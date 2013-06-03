/*global can,document,window */

(function () {
    'use strict';

    var Message, Flipboard;

    Message = can.Observe({
        comment: function () { return this.attr('comment'); },
        author:  function () { return this.attr('author');  },
    });

    Flipboard = can.Control({
        defaults: {
            showEffect: 'flipInX',
            hideEffect: 'flipOutX',
            backgroundColors: [
                '#EC5f41',
                '#24B0CF',
                '#8A69B3',
                '#85B206',
            ],
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
                addClass(this.options.hideEffect).
                removeClass(this.options.showEffect);
        },

        commentWillFlipIn: function () {

            // show the author; it's been updated
            this.author.
                addClass(this.options.showEffect).
                removeClass(this.options.hideEffect);

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
                removeClass(this.options.showEffect).
                addClass(this.options.hideEffect);
        },

        commentDidFlipIn: function () { },

        commentDidFlipOut: function () {

            // sync comment and author
            this.comment.text(this.message.comment());
            this.author.text(this.message.author());

            // finish flipping
            this.comment.
                addClass(this.options.showEffect).
                removeClass(this.options.hideEffect);

            this.flipping = false;
        },

        '.comment webkitAnimationStart': function (element, event) {
            event = event.originalEvent;
            switch (event.animationName) {
            case this.options.hideEffect:
                this.commentWillFlipOut();
                break;
            case this.options.showEffect:
                this.commentWillFlipIn();
                break;
            }
        },

        '.comment webkitAnimationEnd': function (element, event) {
            event = event.originalEvent;
            switch (event.animationName) {
            case this.options.hideEffect:
                this.commentDidFlipOut();
                break;
            case this.options.showEffect:
                this.commentDidFlipIn();
                break;
            }
        }
    });

    function buildEventSource(flipboard) {
        var es = new window.EventSource("/message/events");

        //  todo message-blocked
        //  es.addEventListener("message-blocked", function (event) { });

        es.addEventListener("message-cycled", function (event) {
            var data = JSON.parse(event.data);
            flipboard.message.attr(data);
        });

        return es;
    }


    window.flipboard = new Flipboard('#flipboard', {
        showEffect: 'bounceInRight',
        hideEffect: 'bounceOutRight',
    });

    window.eventSource = buildEventSource(window.flipboard);

}());
