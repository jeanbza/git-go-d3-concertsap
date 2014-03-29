define(['jquery', 'jqueryui'], function ($) {
    $.widget('cs.parser', {
        _init: function() {
            var self = this;
            self.options.elem = $(self.element);

            self.options.elem.keyup(function() {
                var rawText = $(this).val();
                var matches = self.regexMatch(/((([a-zA-Z0-9&\/!]+) ?)+)/g, rawText);
                var parsedString = "";

                matches.forEach(function(item) {
                    parsedString += item+"<br>";
                });

                self.options.elemDisplay.html(parsedString);
                self.options.elemOutput.val(parsedString);
            });
        },

        regexMatch: function(pattern, str) {
            var regex = new RegExp(pattern);
            var matches = [];

            while ((myArray = regex.exec(str)) != null) {
                matches.push(myArray[1]);
            }

            return matches;
        }
    });
});