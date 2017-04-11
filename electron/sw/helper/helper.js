/**
 * Add Listener 
 *          Add listener to the SW
 * @param {Object} bindVar
 * @param {String} event
 * @param {Function} callback 
 */
const addListener = (bindVar, event, callback) => {
    if (callback === null || callback === undefined)
        callback = function(e){console.log(e)};
    
    if (typeof(event) !== 'string')
        throw new Error('event is not a type String');

    bindVar.addListener(event, callback);
};