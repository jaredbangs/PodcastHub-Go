var AppModel = require('./models/app.js');

var appView2 = require('./views/app.js');

$( document ).ready(function() {
	var model = new AppModel();

	var view = new appView2({ model: model });
	view.render();
});
