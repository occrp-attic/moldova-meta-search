(function() {
  'use strict';

  angular
    .module('courtApp')
    .controller('MainController', MainController);

  /** @ngInject */
  function MainController($log) {
    var vm = this;

    vm.results = [];
    vm.keyword = '';


    vm.search = search;

    function search() {
      console.log(vm.keyword);

      vm.results = [1,2,3];
    }

  }
})();
