import * as React from 'react';
import * as _ from 'lodash-es';
import * as PropTypes from 'prop-types';

import { modelFor, k8sPatch } from '../module/k8s';
import { NameValueEditor, NAME, VALUE } from './utils/name-value-editor';
import { PromiseComponent } from './utils';

/**
 * Set up initial value for the environment vars state. Use this in constructor or cancelChanges.
 *
 * @param initialPairObjects
 * @returns {*}
 * @private
 */
const getPairsFromObject = (element) => {
  if (_.isUndefined(element.env)) {
    return [['', '']];
  }
  return _.map(element.env, (leafNode) => {
    return Object.values(leafNode);
  });
};

/**
 * Get name/value pairs from an array or object source
 *
 * @param initialPairObjects
 * @returns {Array}
 */
export const envVarsToArray = (initialPairObjects) => {
  if (_.isArray(initialPairObjects)) {
    return _.map(initialPairObjects, (element) => {
      return getPairsFromObject(element);
    });
  }
  return [getPairsFromObject(initialPairObjects)];
};

export class EnvironmentPage extends PromiseComponent {
  /**
   * Set initial state and decide which kind of env we are setting up
   *
   * @param props
   */
  constructor(props) {
    super(props);

    this.clearChanges = () => this._clearChanges();
    this.saveChanges = (...args) => this._saveChanges(...args);
    this.updateEnvVars = (...args) => this._updateEnvVars(...args);

    const currentEnvVars = envVarsToArray(this.props.rawEnvData);
    this.state = {
      currentEnvVars,
      success: null
    };
  }

  /**
   * Return env var pairs in name value notation, and strip out any pairs that have empty NAME values.
   *
   *
   * @param finalEnvPairs
   * @returns {Array}
   * @private
   */
  _envVarsToNameVal(finalEnvPairs) {
    return _.filter(finalEnvPairs, finalEnvPair => !_.isEmpty(finalEnvPair[NAME]))
      .map(finalPairForContainer => {
        if(finalPairForContainer[VALUE] instanceof Object) {
          return {
            'name': finalPairForContainer[NAME],
            'valueFrom': finalPairForContainer[VALUE]
          };
        }
        return {
          'name': finalPairForContainer[NAME],
          'value': finalPairForContainer[VALUE]
        };
      });
  }

  /**
   * Callback for NVEditor update our state with new values
   * @param env
   * @param i
   */
  _updateEnvVars(env, i=0) {
    const {rawEnvData} = this.props;
    const {currentEnvVars} = this.state;
    const currentEnv = currentEnvVars;
    currentEnv[i] = env.nameValuePairs;
    const modified = !_.isEqual(currentEnv, envVarsToArray(rawEnvData));
    this.setState({
      currentEnvVars: currentEnv,
      success: null,
      modified,
    });
  }

  /**
   * Reset the page to initial state
   * @private
   */
  _clearChanges() {
    const {rawEnvData} = this.props;
    this.setState({
      currentEnvVars: envVarsToArray(rawEnvData),
      errorMessage: null,
      success: null,
      modified: false
    });
  }

  /**
   * Make it so. Patch the values for the env var changes made on the page.
   * 1. Validate for dup keys
   * 2. Throw out empty rows
   * 3. Use add command if we are adding new env vars, and replace if we are modifying
   * 4. Send the patch command down to REST, and update with response
   *
   * @param e
   */
  _saveChanges(e) {
    const {envPath, rawEnvData, obj} = this.props;
    const {currentEnvVars} = this.state;
    e.preventDefault();

    // Convert any blank values to null
    const kind = modelFor(obj.kind);

    const patch = currentEnvVars.map((finalPairsForContainer, i) => {
      let op = 'add';
      const path = _.isArray(rawEnvData) ? `/${envPath.join('/')}/${i}/env` : `/${envPath.join('/')}/env`;
      if (_.isArray(rawEnvData)) {
        if (rawEnvData[i].env) {
          op = 'replace';
        }
      } else {
        if (rawEnvData.env) {
          op = 'replace';
        }
      }
      return {path, op, value: this._envVarsToNameVal(finalPairsForContainer)};
    });

    const promise = k8sPatch(kind, obj, patch);
    this.handlePromise(promise).then((res) => {
      const newEnvData = _.get(res, envPath);
      this.setState({
        success: 'Successfully updated the environment variables.',
        errorMessage: null,
        currentEnvVars: envVarsToArray(newEnvData),
        rawEnvData: newEnvData,
        modified: false
      });
    });
  }

  render() {
    const {errorMessage, success, inProgress, currentEnvVars} = this.state;
    const {rawEnvData, readOnly} = this.props;

    const containerVars = currentEnvVars.map((envVar, i) => {
      const keyString = _.isArray(rawEnvData) ? rawEnvData[i].name : rawEnvData.from.name;
      return <div key={keyString} className="co-m-pane__body-group">
        { _.isArray(rawEnvData) && <h1 className="co-section-title">Container {keyString}</h1> }
        <NameValueEditor nameValueId={i} nameValuePairs={envVar} updateParentData={this.updateEnvVars} addString="Add Value" nameString="Name" readOnly={readOnly}/>
      </div>;
    });

    return <div className="co-m-pane__body">
      {containerVars}
      <div className="co-m-pane__body-group">
        <div className="environment-buttons">
          {errorMessage && <p className="alert alert-danger"><span className="pficon pficon-error-circle-o"></span>{errorMessage}</p>}
          {success && <p className="alert alert-success"><span className="pficon pficon-ok"></span>{success}</p>}
          {!readOnly && <button disabled={inProgress} type="submit" className="btn btn-primary" onClick={this.saveChanges}>Save Changes</button>}
          {this.state.modified && <button type="button" className="btn btn-link" onClick={this.clearChanges}>Clear Changes</button>}
        </div>
      </div>
    </div>;
  }
}
EnvironmentPage.propTypes = {
  obj: PropTypes.object.isRequired,
  rawEnvData: PropTypes.oneOfType([PropTypes.object, PropTypes.array]).isRequired,
  envPath: PropTypes.array.isRequired,
  readOnly: PropTypes.bool.isRequired
};
