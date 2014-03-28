require.config({
    baseUrl: '/static/js/',
    urlArgs: "bust=" + (new Date()).getTime(),
    paths: {
        'jquery': 'external/jquery/jquery-2.0.3.min',
        'jqueryui': 'external/jquery-ui/jquery-ui-1.10.4.custom',
        'bootstrap': 'external/bootstrap/bootstrap.min',
        'bootstrap-dtp': 'external/bootstrap/bootstrap-datetimepicker.min'
    },

    shim: {  
        "jqueryui": {
            deps: ['jquery']
        },
        "bootstrap" : {
            deps: ['jquery']
        },
        'bootstrap-dtp': {
            deps: ['jquery', 'bootstrap']
        }
    }
});

require([
    'jquery',
    'bootstrap',
    'modules/parser'
], 
function ($) {
    $(document).ready(function() {
        $('.dropdown-toggle').dropdown();

        $(".cs-js-parse-input").parser({
            elemDisplay: $(".cs-js-parse-display"),
            elemOutput: $(".cs-js-parse-output"),
        });
    });
});