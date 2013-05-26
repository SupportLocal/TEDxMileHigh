/*global can,document,window */

(function () {
    'use strict';

    var Flipboard = can.Control({
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

        init: function () {
            this.comment = this.element.find('.comment');
            this.author = this.element.find('.author');

            this.backgroundColor = can.compute(this.options.backgroundColors[0]);
            this.backgroundColor.bind('change', function (event, newColor, oldColor) {
                document.body.style.backgroundColor = newColor;
            });
        },

        flip: function () {
            var flipIn = this.options.flipIn,
                flipOut = this.options.flipOut;

            if (this.comment.is('.' + flipOut)) {
                this.comment.addClass(flipIn).removeClass(flipOut);
            } else {
                this.comment.addClass(flipOut).removeClass(flipIn);
            }
        },

        commentWillFlipIn: function () {
            // show the author; it's been updated
            this.author.show();

            // compute and set new background color
            var backgroundColors = this.options.backgroundColors,
                curIdx = backgroundColors.indexOf(this.backgroundColor()),
                nextIdx = ((curIdx + 1) % backgroundColors.length);
            this.backgroundColor(backgroundColors[nextIdx]);
        },

        commentWillFlipOut: function () {
            // hide the author; it will be updated
            this.author.hide();
        },

        commentDidFlipIn: function () { },

        commentDidFlipOut: function () {
            // update comment
            // update author
            this.author.text('@levicook');
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
