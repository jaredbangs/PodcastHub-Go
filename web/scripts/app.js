var AppController = require('./controllers/app.js');
var AppLayoutView = require('./views/appLayout.js');
var AppRouter = require('./routers/app.js');

module.exports = Marionette.Application.extend({

	initialize: function () {
		this.controller = new AppController({ application: this });
		this.router = new AppRouter({ controller: this.controller });
	},

	onBeforeStart: function () {
		this.layout = new AppLayoutView();
		this.layout.render();
	},

	onStart: function () {
		Backbone.history.start();
		
		this.controller.showFeedList();
	}
});
