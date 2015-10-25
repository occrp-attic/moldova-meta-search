(function() {
  'use strict';

  describe('controllers', function(){
    var vm;

    beforeEach(module('courtApp'));
    beforeEach(inject(function(_$controller_) {

      vm = _$controller_('MainController');
    }));

    it('should define more than 5 awesome things', function() {
      expect(angular.isArray(vm.awesomeThings)).toBeTruthy();
      expect(vm.awesomeThings.length === 3).toBeTruthy();
    });
  });
})();
