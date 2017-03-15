const assert = require('assert');

const myApp = {
    name : 'colorPicker',
    author: 'Marc'
};

describe('myApp', () => {
    it('should return true', () => {
        assert.equal('colorPicker', myApp.name);
    });
});