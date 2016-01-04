var lmapp = angular.module('lmjrfll', [], function($interpolateProvider) {
  $interpolateProvider.startSymbol('[[');
  $interpolateProvider.endSymbol(']]');
});

lmapp.controller('HeaderController', function($scope, $http) {
  $http.get('/api/1/expo/current').success(function(data){
    $scope.expo = data;
  });

});
