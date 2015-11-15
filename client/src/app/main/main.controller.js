(function () {
  'use strict';

  angular
    .module('courtApp')
    .controller('MainController', MainController)
    .filter('checkTypes', function() {
      return function(decisions, types) {
        var filteredDecisions = [];

        if (types.length) {
          angular.forEach(decisions, function (decision) {
            if (types.indexOf(decision.type) > -1) {
              filteredDecisions.push(decision);
            }
          });
        } else {
          filteredDecisions = decisions;
        }

        return filteredDecisions;
      };
    })
    .filter('checkCourts', function() {
      return function(decisions, courts) {
        var filteredDecisions = [];

        if (courts.length) {
          angular.forEach(decisions, function (decision) {
            if (courts.indexOf(decision.court) > -1) {
              filteredDecisions.push(decision);
            }
          });
        } else {
          filteredDecisions = decisions;
        }

        return filteredDecisions;
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

    vm.selectedTypes = [];
    vm.selectedCourts = [];

    function search() {
      if (vm.keyword.length <= 0) {
          return;
      }
      vm.loading = true;
      $http.get('api/v1/courts/search/' + vm.keyword).success(function(data) {
        var courts = {},
            types = {};

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
        vm.results = data;
        vm.loading = false;
      });
    }

    // toggle selection for a given employee by name
    vm.toggleSelection = function toggleSelection(type, collection) {
      var selectedIDX = collection.indexOf(type);

      if (selectedIDX > -1) {
        collection.splice(selectedIDX, 1);
      } else {
        collection.push(type);
      }
    };
  }
})();
