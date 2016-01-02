var lmapp = angular.module('lmjrfll', [], function($interpolateProvider) {
  $interpolateProvider.startSymbol('[[');
  $interpolateProvider.endSymbol(']]');
});

lmapp.controller('HeaderController', function() { });
