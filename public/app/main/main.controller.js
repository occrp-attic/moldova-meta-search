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
    vm.loading = false;

    vm.search = search;

    function search() {
      vm.loading = true;
      $http.get('api/v1/courts/search/' + vm.keyword).success(function(data) {
        vm.results = data;
        vm.loading = false;
      });
    }

  }
})();
