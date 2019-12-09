var display = document.getElementById('display');
var error   = document.getElementById('error');
var buttons = document.getElementsByClassName('btn');
var calc    = '';

function update(disp, err = '') {
    display.innerText = disp;
    error.innerText = err;
}

for(var i = 0; i < buttons.length; i++)
{
  buttons[i].onclick= function(text) {
    return function() {

      if (text == '=') {

        Calculator.Evaluate(calc).then((reply) => {

            update(calc = reply.Result);

        }).catch((err) => {

            update(calc, err);

        })

      } else if (text == 'C'){

        update(calc = '');

      } else {

        update(calc += (text.length > 2 ? text+'(' : text));

      }
        
    }
  }(buttons[i].innerText);
}
