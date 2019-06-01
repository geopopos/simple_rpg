import 'phaser';

export default class BootScene extends Phaser.Scene {
    constructor(key){
        super(key);
    }

    preload() {
        this.levels = {
            1: 'level1',
            2: 'level2'
        };
        // load in the tilemap
        this.load.tilemapTiledJSON('level1','../../assets/tilemaps/level1.json');
        this.load.tilemapTiledJSON('level2','../../assets/tilemaps/level2.json');
        // load in the spritesheet
        this.load.spritesheet('RPGpack_sheet', '../../assets/spritesheets/RPGpack_sheet.png', {frameWidth: 64, frameHeight: 64});
        this.load.spritesheet('characters', '../../assets/spritesheets/personajes-lanto.png', {frameWidth: 32, frameHeight: 32});
        this.load.image('portal', '../../assets/spritesheets/raft.png', {frameWidth: 64, frameHeight: 64});
        // load in our coin sprite
        this.load.image('coin', '../../assets/spritesheets/coin_01.png');
    }

    create() {
        this.scene.start('Game', {level: 1, newGame: true, levels: this.levels });
    }
}