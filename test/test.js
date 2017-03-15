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

describe('author', () => {
    it('should return true for author', () => {
        assert.equal('Marc', myApp.author);
    });
});