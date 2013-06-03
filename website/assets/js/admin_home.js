/*global can,document,window */

(function () {
    'use strict';

    var dataPool, CurrentController, PendingController, Router;

    dataPool = document.getElementById('data-pool');

    CurrentController = can.Control({
        defaults: { view: '/ejs/admin_home/current-message.ejs' }
    }, {

        init: function (element, options) {
            this.message = new can.Observe(options.message);
            element.append(can.view(options.view, {
                message: this.message
            }));
        },

    });

    PendingController = can.Control({
        defaults: { view: '/ejs/admin_home/pending-messages.ejs' }
    }, {

        init: function (element, options) {
            this.messages = new can.Observe.List(options.messages);
            element.append(can.view(options.view, {
                messages: this.messages
            }));
        },

        'a click': function (link, event) {
            event.preventDefault();

            var message = link.parent('.message').data('message');
            console.log(message);
        }

    });


    Router = can.Control({
        defaults: {
            view: '/ejs/admin_home.ejs',
            currentContainer: '#current-message',
            pendingContainer: '#pending-messages',
        }
    }, {

        init: function (element, options) {
            element.append(can.view(options.view));

            this.currentController = new CurrentController(options.currentContainer, {
                message: options.current,
            });

            this.pendingController = new PendingController(options.pendingContainer, {
                messages: options.messages,
            });
        }

    });


    function buildEventSource(router) {
        var es = new window.EventSource("/message/events"),
            currentController = router.currentController;

        es.onmessage = function (event) {
            var data = JSON.parse(event.data);
            currentController.message.attr(data);
        };

        return es;
    }

    // bind our globals ---

    window.data = dataPool ? JSON.parse(dataPool.text) : {};

    window.router = new Router(document.body, {
        current: window.data.current,
        messages: window.data.messages,
    });

    window.eventSource = buildEventSource(window.router);

}());
