var lmapp = angular.module('lmjrfll', [], function($interpolateProvider) {
  $interpolateProvider.startSymbol('[[');
  $interpolateProvider.endSymbol(']]');
});

lmapp.controller('HeaderController', ['$scope', '$http', function($scope, $http) {
  $http.get('/api/1/user/profile').success(function(data){
    $scope.user = data;
  });
  $http.get('/api/1/expo/current').success(function(data){
    $scope.expo = data;
  });
}]);
