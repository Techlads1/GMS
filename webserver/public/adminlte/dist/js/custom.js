$(function () {
    $('[data-mask]').inputmask()
})

$(document).ready(function() {
$(".check-staff").click(function(){
      $("#staff").prop("checked", true);
      $("#year").hide();
      $("#csee_number").hide();
      $("#check_number").show();
      $("#check_number_id").prop('required',true);
      $("#csee_number_id").prop('required',false);
      $("#year_id").prop('required',false);
      
  });
  $(".check-student").click(function(){
      $("#student").prop("checked", true);
      $("#year").show();
      $("#csee_number").show();
      $("#check_number").hide();
      $("#csee_number_id").prop('required',true);
      $("#year_id").prop('required',true);
      $("#check_number_id").prop('required',false);
  });

$('#password').keyup(function() {
    var password = $('#password').val();
    if (checkStrength(password)) {
        $('#sign-up').attr('disabled', false);
    }else{
        $('#sign-up').attr('disabled', true);
    }
});

$('#confirm-password').blur(function() {
    if ($('#password').val() !== $('#confirm-password').val()) {
        $('#popover-cpassword').show().text("Password Mismatch");
        $('#sign-up').attr('disabled', true);
    } else {
        if (checkStrength($('#password').val())) {
        $('#popover-cpassword').hide();
        $('#sign-up').attr('disabled', false);
        }
    }
});

function checkStrength(password) {
    var strength = 0;

    //If password contains both lower and uppercase characters, increase strength value.
    if (password.match(/([a-z].*[A-Z])|([A-Z].*[a-z])/)) strength += 1;
    //If it has numbers and characters, increase strength value.
    if (password.match(/([a-zA-Z])/) && password.match(/([0-9])/))  strength += 1
    //If it has one special character, increase strength value.
    if (password.match(/([!,%,&,@,#,$,^,*,?,_,~])/))  strength += 1
    if (password.length > 8)  strength += 1
  
    // If value is less than 2
    if (strength < 2) {
        $('#result').removeClass()
        $('#password-strength').addClass('bg-danger');
        $('#result').addClass('text-danger').text('Very Week');
        $('#password-strength').css('width', '10%');
        return false;
    }
    if (strength === 2) {
        $('#result').addClass('good');
        $('#password-strength').removeClass('bg-danger');
        $('#password-strength').addClass('bg-warning');
        $('#result').addClass('text-warning').text('Moderate')
        $('#password-strength').css('width', '60%');
        return false;
    }
    if (strength === 4) {
        $('#result').removeClass()
        $('#result').addClass('strong');
        $('#password-strength').removeClass('bg-danger','bg-warning');
        $('#password-strength').addClass('bg-success');
        $('#result').addClass('text-success').text('Strong');
        $('#password-strength').css('width', '100%');
        return true;
    }
 }

});