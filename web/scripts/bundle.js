(function e(t,n,r){function s(o,u){if(!n[o]){if(!t[o]){var a=typeof require=="function"&&require;if(!u&&a)return a(o,!0);if(i)return i(o,!0);var f=new Error("Cannot find module '"+o+"'");throw f.code="MODULE_NOT_FOUND",f}var l=n[o]={exports:{}};t[o][0].call(l.exports,function(e){var n=t[o][1][e];return s(n?n:e)},l,l.exports,e,t,n,r)}return n[o].exports}var i=typeof require=="function"&&require;for(var o=0;o<r.length;o++)s(r[o]);return s})({1:[function(require,module,exports){
var AppModel = require('./models/app.js');

var appView = require('./views/app.js');

$( document ).ready(function() {
	var model = new AppModel();

	var view = new appView({ model: model });
	view.render();
});

},{"./models/app.js":2,"./views/app.js":3}],2:[function(require,module,exports){
module.exports = AppModel = Backbone.Model.extend({



});


},{}],3:[function(require,module,exports){
require('../models/app.js')

module.exports = AppView = Backbone.View.extend({

	render: function() {
		console.log("Hey");
	}

});

},{"../models/app.js":2}]},{},[1]);
