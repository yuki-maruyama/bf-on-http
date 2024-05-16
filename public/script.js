document.getElementById('inputForm').addEventListener('submit', async function(event) {
  event.preventDefault();
  document.getElementById('outputDisplay').textContent = "Wait ..."
  const userInput = document.getElementById('scriptForm').value;
  const req = new Request('/run', {method: "POST", body: userInput})
  fetch(req)
    .then(async(res) => {
      if (res.status != 200){
        document.getElementById('outputDisplay').style.color = "red"
      } else {
        document.getElementById('outputDisplay').style.color = "black"
      }
      document.getElementById('outputDisplay').textContent = await res.text()
    })
});
