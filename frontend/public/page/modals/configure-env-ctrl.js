angular.module('app')
.controller('ConfigureEnvCtrl', function(_, $scope, $modalInstance, $controller,
      $rootScope, container, k8s) {
  'use strict';

  $scope.rowMgr = $controller('RowMgr', {
    $scope: $rootScope.$new(),
    emptyCheck: function(e) {
      return _.isEmpty(e.name) || _.isEmpty(e.value);
    },
    getEmptyItem: k8s.docker.getEmptyEnvVar,
  });

  $scope.initEnvVars = function(envVars) {
    if (_.isEmpty(envVars)) {
      $scope.rowMgr.setItems([]);
    } else {
      $scope.rowMgr.setItems(angular.copy(envVars));
    }
  };

  $scope.save = function() {
    container.env = $scope.rowMgr.getNonEmptyItems();
    $modalInstance.close(container);
  };

  $scope.cancel = function() {
    $modalInstance.dismiss('cancel');
  };

  $scope.initEnvVars(container.env);
})
.controller('ConfigureEnvFormCtrl', function($scope) {
  'use strict';
  $scope.submit = $scope.save;
});