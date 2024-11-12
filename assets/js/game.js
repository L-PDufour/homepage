const canvas = document.getElementById("canvas");
const ctx = canvas.getContext("2d");
canvas.width = 1024
canvas.height = 576

class Sprite {
  constructor({ position }) {
    this.position = position
  }

  draw() {
    ctx.fillStyle = 'red'
    ctx.fillRect(this.position.x, this.position.y, 10, 50,);
  }
  update() {
    console.log(this.position)
    this.draw()
  }
}
class Wall {
  constructor({ position, blockCount }) {
    this.blocks = [];
    this.position = position;
    this.createBlocks(blockCount);
  }

  createBlocks(count) {
    for (let i = 0; i < count; i++) {
      this.blocks.push(
        new Block({
          position: {
            x: this.position.x,
            y: this.position.y + i * 50,
          }
        })
      );
    }
  }

  update() {
    this.blocks.forEach((block) => block.update());
  }
}

class Block extends Sprite {
  constructor({ position }) {
    super({ position });
    this.direction = 1;
    this.speed = Math.floor(Math.random() * 10) + 1;
  }

  update() {
    this.draw();

    this.position.x += this.speed * this.direction;

    if (this.position.x >= canvas.width - 10) {
      this.direction = -1;
    } else if (this.position.x <= 0) {
      this.direction = 1;
    }
  }
}



function handleKeyDown(e) {
  switch (e.key) {
    case "ArrowUp":
      player.position.y -= 10;
      break;
    case "ArrowDown":
      player.position.y += 10;
      break;
    case "ArrowLeft":
      player.position.x -= 10;
      break;
    case "ArrowRight":
      player.position.x += 10;
      break;
  }
}


const screen = {
  draw() {
    ctx.clearRect(0, 0, canvas.width, canvas.height);
    ctx.fillStyle = "white";
    ctx.fillRect(0, 0, canvas.width, canvas.height);
  },
};

document.addEventListener("keydown", handleKeyDown);
const block = {
  x: canvas.width,
  y: canvas.height / 2,
  draw() {
    ctx.fillStyle = "red";
    ctx.fillRect(this.x, this.y, this.x + 2, this.y - 10);
    this.x -= 1
  }
}

const player = new Sprite({ position: { x: 0, y: 0 } })
const wall = new Wall({
  position: { x: canvas.width - 10, y: 0 },
  blockCount: 15
});

function animate() {
  window.requestAnimationFrame(animate);
  screen.draw();
  player.update()
  wall.update()
}
animate();
