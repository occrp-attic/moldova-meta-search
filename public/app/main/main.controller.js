(function () {
  'use strict';

  angular
    .module('courtApp')
    .controller('MainController', MainController)
    .filter('check', function() {
      return function(decisions, types) {
        console.log(decisions, types);

        var filtereDecisions = [];

        if (types.length) {
          angular.forEach(decisions, function (decision) {
            if (types.indexOf(decision.type) > -1) {
              filtereDecisions.push(decision);
            }
          });

          return filtereDecisions;
        } else {
          return decisions;
        }
      };
    })
    .filter('checkCourts', function() {
      return function(decisions, courts) {
        console.log(decisions, courts);

        var filtereDecisions = [];

        if (courts.length) {
          angular.forEach(decisions, function (decision) {
            if (courts.indexOf(decision.court) > -1) {
              filtereDecisions.push(decision);
            }
          });

          return filtereDecisions;
        } else {
          return decisions;
        }
      };
    });

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
      //$http.get('data/data.json').success(function (data) {
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
        $log.log(types);
        $log.log(courts);
        vm.results = data;
        vm.loading = false;
      });
    }

    vm.selection = [];
    vm.selectedCourts = [];
    // toggle selection for a given employee by name
    vm.toggleSelection = function toggleSelection(type) {
      var idx = vm.selection.indexOf(type);

      // is currently selected
      if (idx > -1) {
        vm.selection.splice(idx, 1);
      }

      // is newly selected
      else {
        vm.selection.push(type);
      }
    };

    vm.toggleSelectedCourts = function toggleSelectedCourts(type) {
      var idx = vm.selectedCourts.indexOf(type);

      // is currently selected
      if (idx > -1) {
        vm.selectedCourts.splice(idx, 1);
      }

      // is newly selected
      else {
        vm.selectedCourts.push(type);
      }
    };
  }

})();
