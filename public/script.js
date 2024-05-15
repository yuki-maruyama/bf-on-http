document.getElementById('inputForm').addEventListener('submit', async function(event) {
  event.preventDefault();
  const userInput = document.getElementById('scriptForm').value;
  const req = new Request('/run', {method: "POST", body: userInput})
  const res = await fetch(req)
  document.getElementById('outputDisplay').textContent = await res.text()
});
