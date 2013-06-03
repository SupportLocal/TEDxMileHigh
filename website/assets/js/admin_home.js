/*global can,document,window */

(function () {
    'use strict';

    var CurrentController, PendingController, Router;

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

        'a.block click': function (link, event) {
            event.preventDefault();

            var messageEl = link.parent('.message'),
                message = messageEl.data('message'),
                json = { ids: [message.attr('id')] };

            can.ajax({
                url: '/admin/messages/block',
                data: JSON.stringify(json),
                type: 'POST'
            });
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
            var router = this;

            element.append(can.view(options.view));

            router.currentController = new CurrentController(options.currentContainer, { message:  options.current, });
            router.pendingController = new PendingController(options.pendingContainer, { messages: options.messages });

            router.eventSource = new window.EventSource("/message/events");

            router.eventSource.addEventListener("message-blocked", function (event) {
                router.messageBlocked(JSON.parse(event.data));
            });

            router.eventSource.addEventListener("message-cycled", function (event) {
                router.messageCycled(JSON.parse(event.data));
            });
        },

        messageBlocked: function (message) {
            // TODO
            console.log({ messageBlocked: message });
        },

        messageCycled: function (message) {
            this.currentController.message.attr(message);
        },

    });

    // bind our globals ---

    window.data = JSON.parse(document.getElementById('data-pool').text);

    window.router = new Router(document.body, {
        current: window.data.current,
        messages: window.data.messages,
    });

}());
