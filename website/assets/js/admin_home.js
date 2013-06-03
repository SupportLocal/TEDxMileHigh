/*global can,document,window */

(function () {
    'use strict';

    var dataPool, ListController, Router;

    dataPool = document.getElementById('data-pool');

    ListController = can.Control({
        defaults: {
            view: '/ejs/admin_home/list.ejs',
        }
    }, {

        init: function (element, options) {
            this.messages = new can.Observe.List(options.messages);

            element.append(can.view(options.view, {
                messages: this.messages
            }));
        },

    });


    Router = can.Control({
        defaults: {
            view: '/ejs/admin_home.ejs',
            listContainer: '#message-list',
        }
    }, {

        init: function (element, options) {
            element.append(can.view(options.view));

            this.listController = new ListController(options.listContainer, {
                messages: options.data.messages,
            });
        }

    });

    window.data = dataPool ? JSON.parse(dataPool.text) : {};
    window.router = new Router(document.body, { data: window.data });

}());
