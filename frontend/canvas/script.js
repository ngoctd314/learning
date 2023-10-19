function draw() {
  const canvas = document.getElementById("root");
  if (canvas.getContext) {
    const ctx = canvas.getContext("2d");

    ctx.beginPath();
    ctx.moveTo(0, 0);
    ctx.lineTo(100, 50);
    ctx.lineTo(90, 30);
    ctx.lineTo(90, 50);
    ctx.fill();

    // ctx.fillStyle = "rgba(0, 0, 200, 0.5)";
    // ctx.fillRect(30, 30, 50, 50);
  }
}
window.addEventListener("load", draw);
