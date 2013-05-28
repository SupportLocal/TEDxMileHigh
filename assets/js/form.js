/*global can,document,window */

(function () {
    'use strict';

    var InboundMessage, FormController, Router;

    InboundMessage = can.Observe({

    });

    FormController = can.Control({
        defaults: {
            view: '/ejs/form.ejs'
        }
    }, {

        init: function (element, options) {
            this.inboundMessage = new InboundMessage();
            element.append(can.view(options.view, this.inboundMessage));
        },

        'input[type=text],textarea change': function (input) {
            this.inboundMessage.attr(
                input.prop('id'),
                input.val()
            );
        },

        'form submit': function (form, event) {
            event.preventDefault();
            console.log(form);

            var inboundMessage = this.inboundMessage;

            can.ajax({
                url: '/',
                data: JSON.stringify(inboundMessage.serialize()),
                type: 'POST'
            }).then(function (response) {
                inboundMessage.attr(response.d);
            });
        }

    });

    Router = can.Control({
        defaults: {
            container: '#container'
        }
    }, {

        init: function (element, options) {
            this.formController = new FormController(options.container);
        }

    });

    window.router = new Router(document.body);

}());
