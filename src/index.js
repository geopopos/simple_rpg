import Phaser from "phaser";
import config from "./config";
import GameScene from "./scenes/Game"
import BootScene from "./scenes/Boot"

class Game extends Phaser.Game {
  constructor() {
    super(config);
    this.scene.add('Boot', BootScene);
    this.scene.add('Game', GameScene);
    this.scene.start('Boot');
  }
}

window.game = new Game();
window.addEventListener('resize', (event) => {
  window.game.scale.resize(window.innerWidth, window.innerHeight);
});