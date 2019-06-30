$("#myForm").submit(function(e) {

    e.preventDefault(); // avoid to execute the actual submit of the form.

    var form = $(this);
    var url = form.attr('action');

    $.ajax({
        type: "GET",
        url: url,
        data: form.serialize(), // serializes the form's elements.
        success: function(data)
        {
            alert(data); // show response from the php script.
        }
    });


});
