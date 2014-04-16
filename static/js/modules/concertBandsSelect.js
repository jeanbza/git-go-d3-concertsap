define(['jquery', 'jqueryui'], function ($) {
    $.widget('cs.concertBandsSelect', {
        _init: function() {
            var self = this;
            self.options.elem = $(self.element);
            self.options.concertSelectElem = $("."+self.options.elem.attr("concert-select-class"));
            
            self.updateConcertBands(self.options.concertSelectElem.val());

            self.options.concertSelectElem.change(function() {
                self.updateConcertBands(self.options.concertSelectElem.val());
            });
        },

        updateConcertBands: function(concertId) {
            var self = this;
            
            self.options.elem.empty();

            $.ajax({
                url: "/api/getConcertBands/"+concertId,
                success: function(data) {
                    var bands = JSON.parse(data);

                    bands.forEach(function(band) {
                        self.options.elem.append("<option value='"+band.BandId+"'>"+band.BandName+"</option>");
                    });
                }
            });
        }
    });
});