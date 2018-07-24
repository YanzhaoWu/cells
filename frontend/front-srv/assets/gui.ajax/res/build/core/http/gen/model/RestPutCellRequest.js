/**
 * Pydio Cells Rest API
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: 1.0
 * 
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 * Do not edit the class manually.
 *
 */

'use strict';

exports.__esModule = true;

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { 'default': obj }; }

function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError('Cannot call a class as a function'); } }

var _ApiClient = require('../ApiClient');

var _ApiClient2 = _interopRequireDefault(_ApiClient);

var _RestCell = require('./RestCell');

var _RestCell2 = _interopRequireDefault(_RestCell);

/**
* The RestPutCellRequest model module.
* @module model/RestPutCellRequest
* @version 1.0
*/

var RestPutCellRequest = (function () {
    /**
    * Constructs a new <code>RestPutCellRequest</code>.
    * @alias module:model/RestPutCellRequest
    * @class
    */

    function RestPutCellRequest() {
        _classCallCheck(this, RestPutCellRequest);

        this.Room = undefined;
        this.CreateEmptyRoot = undefined;
    }

    /**
    * Constructs a <code>RestPutCellRequest</code> from a plain JavaScript object, optionally creating a new instance.
    * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
    * @param {Object} data The plain JavaScript object bearing properties of interest.
    * @param {module:model/RestPutCellRequest} obj Optional instance to populate.
    * @return {module:model/RestPutCellRequest} The populated <code>RestPutCellRequest</code> instance.
    */

    RestPutCellRequest.constructFromObject = function constructFromObject(data, obj) {
        if (data) {
            obj = obj || new RestPutCellRequest();

            if (data.hasOwnProperty('Room')) {
                obj['Room'] = _RestCell2['default'].constructFromObject(data['Room']);
            }
            if (data.hasOwnProperty('CreateEmptyRoot')) {
                obj['CreateEmptyRoot'] = _ApiClient2['default'].convertToType(data['CreateEmptyRoot'], 'Boolean');
            }
        }
        return obj;
    };

    /**
    * @member {module:model/RestCell} Room
    */
    return RestPutCellRequest;
})();

exports['default'] = RestPutCellRequest;
module.exports = exports['default'];

/**
* @member {Boolean} CreateEmptyRoot
*/