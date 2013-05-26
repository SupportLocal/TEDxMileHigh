/*global can,document,window */

(function () {
    'use strict';

    var Flipboard = can.Control({
        defaults: {
            backgroundColors: ['#EC5f41', '#24B0CF', '#8A69B3', '#85B206'],
            flipIn: 'flipInX',
            flipOut: 'flipOutX'
        }
    }, {

        init: function () {
            this.comment = this.element.find('.comment');
            this.author = this.element.find('.author');
            this.flipping = false;

            this.backgroundColor = can.compute(this.options.backgroundColors[0]);
            this.backgroundColor.bind('change', function (event, newColor, oldColor) {
                document.body.style.backgroundColor = newColor;
            });
        },

        flip: function () {
            if (this.flipping) { return; }

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
            // compute and set new message

            /* todo
            var all     = ??,
                curr    = ??,
                currIdx = ??,
                nextIdx = ??,
                next    = ??;
            */

            var next = {
                comment: 'Foo bar bin basz!!',
                author: '@levicook'
            };

            // update comment and author
            this.comment.text(next.comment);
            this.author.text(next.author);

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

}());
