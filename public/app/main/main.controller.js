(function() {
  'use strict';

  angular
    .module('courtApp')
    .controller('MainController', MainController);

  /** @ngInject */
  function MainController($log, $http) {
    var vm = this;

    vm.results = [];
    vm.keyword = '';


    vm.search = search;

    function search() {
      $http.get('data/data.json').success(function(data) {
        vm.results = data;
      });
    }

  }
})();
