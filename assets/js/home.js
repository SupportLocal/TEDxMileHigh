/*global can,document,window */

(function () {
    'use strict';

    var Message, Flipboard, echoSocket;

    Message = can.Observe({
        comment: function () { return this.attr('comment'); },
        author:  function () { return this.attr('author');  },
    });

    Flipboard = can.Control({
        defaults: {
            backgroundColors: [
                '#EC5f41',
                '#24B0CF',
                '#8A69B3',
                '#85B206',
            ],
            flipIn: 'flipInX',
            flipOut: 'flipOutX'
        }
    }, {

        init: function (element, options) {
            var flipboard = this;

            flipboard.flipping = false;

            flipboard.comment = flipboard.element.find('.comment');
            flipboard.author = flipboard.element.find('.author');

            flipboard.message = new Message({
                comment: flipboard.comment.text(),
                author: flipboard.author.text(),
            });

            flipboard.message.bind('change', function () {
                flipboard.flip();
            });

            flipboard.backgroundColor = can.compute(options.backgroundColors[0]);
            flipboard.backgroundColor.bind('change', function (event, newColor) {
                document.body.style.backgroundColor = newColor;
            });
        },

        flip: function () {
            if (this.flipping) {
                return;
            }

            // start flipping
            this.flipping = true;

            this.comment.
                addClass(this.options.flipOut).
                removeClass(this.options.flipIn);
        },

        commentWillFlipIn: function () {

            // show the author; it's been updated
            this.author.
                addClass(this.options.flipIn).
                removeClass(this.options.flipOut);

            // compute and set new background color
            var all     = this.options.backgroundColors,
                curr    = this.backgroundColor(),
                currIdx = all.indexOf(curr),
                nextIdx = (currIdx + 1) % all.length,
                next    = all[nextIdx];

            this.backgroundColor(next);
        },

        commentWillFlipOut: function () {

            // hide the author; it will be updated
            this.author.
                removeClass(this.options.flipIn).
                addClass(this.options.flipOut);
        },

        commentDidFlipIn: function () { },

        commentDidFlipOut: function () {

            // sync comment and author
            this.comment.text(this.message.comment());
            this.author.text(this.message.author());

            // finish flipping
            this.comment.
                addClass(this.options.flipIn).
                removeClass(this.options.flipOut);

            this.flipping = false;
        },

        '.comment webkitAnimationStart': function (element, event) {
            event = event.originalEvent;
            switch (event.animationName) {
            case this.options.flipOut:
                this.commentWillFlipOut();
                break;
            case this.options.flipIn:
                this.commentWillFlipIn();
                break;
            }
        },

        '.comment webkitAnimationEnd': function (element, event) {
            event = event.originalEvent;
            switch (event.animationName) {
            case this.options.flipOut:
                this.commentDidFlipOut();
                break;
            case this.options.flipIn:
                this.commentDidFlipIn();
                break;
            }
        }
    });

    window.flipboard = new Flipboard('#flipboard');


    echoSocket = new window.WebSocket("ws://" + window.location.host + "/echo");

    echoSocket.onopen = function () {
        var heartbeat = 0;
        setInterval(function () {
            heartbeat = heartbeat + 1;
            echoSocket.send("heartbeat: " + heartbeat);
        }, 2500);
    };

    echoSocket.onmessage = function (event) {
        console.log(event.data);
    };

}());
