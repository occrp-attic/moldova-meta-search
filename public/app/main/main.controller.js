(function() {
  'use strict';

  angular
    .module('courtApp')
    .controller('MainController', MainController);

  /** @ngInject */
  function MainController($log, $http) {
    var vm = this;

    vm.results = vm.types = vm.courts = [];
    vm.keyword = '';
    vm.loading = false;

    vm.isCollapsed = false;
    vm.search = search;

    function search() {
      vm.loading = true;
      $http.get('api/v1/courts/search/' + vm.keyword).success(function(data) {
        var courts = {};
        var types = {};

        for (var i = 0; i < data.length; i++) {
          if (data[i] && data[i].court && !courts[data[i].court]) {
            courts[data[i].court] = 1;
          } else {
            courts[data[i].court]++;
          }
          if (data[i] && data[i].type && !types[data[i].type]) {
            types[data[i].type] = 1;
          } else {
            types[data[i].type]++;
          }
        }
        vm.courts = courts;
        vm.types = types;
        console.log(types);
        console.log(courts);
        vm.results = data;
        vm.loading = false;
      });
    }
    
  }
})();
