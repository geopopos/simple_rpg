import 'phaser';
import Player from '../sprites/player';
import Portal from '../sprites/portal';
import Coins from '../groups/coins';
import Enemies from '../groups/enemies';

export default class GameScene extends Phaser.Scene {
    constructor(key){
        super(key);
    }

    init(data) {
        this._LEVEL = data.level;
        this._LEVELS = data.levels;
        this._NEWGAME = data.newGame;
        this.loadingLevel = false;
        if(this._NEWGAME) this.events.emit('newGame');
    }

    create() {
        // listen for the resize event
        this.events.on('resize', this.resize, this);
        // listen for player inpu
        this.cursors = this.input.keyboard.createCursorKeys();

        // create our tilemap
        this.createMap();
        // create our player
        this.createPlayer();
        // creating the portal
        this.createPortal();
        // creating the coins
        this.coins = this.map.createFromObjects('Coins', 'Coin', { key: 'coin' });
        // create group to hold coin sprites
        this.coinsGroup = new Coins(this.physics.world, this, [], this.coins);
        // creating the enemies
        this.enemies = this.map.createFromObjects('Enemies', 'Enemy', {});
        this.enemiesGroup = new Enemies(this.physics.world, this, [], this.enemies);
        
        // add collisions
        this.addCollisions();
        
        // update our camera
        this.cameras.main.startFollow(this.player);

        // create animations for player movement
        this.anims.create({
            key: 'down',
            frames: this.anims.generateFrameNumbers('characters', {frames: [1,2,1,0]}),
            frameRate: 10,
        });
        this.anims.create({
            key: 'left',
            frames: this.anims.generateFrameNumbers('characters', {frames:[13,14,13,12]}),
            frameRate: 5,
        });
        this.anims.create({
            key: 'right',
            frames: this.anims.generateFrameNumbers('characters', {frames:[25,26,25,24]}),
            frameRate: 10,
        });
        this.anims.create({
            key: 'up',
            frames: this.anims.generateFrameNumbers('characters', {frames:[37,38,37,36]}),
            frameRate: 10,
        });
    }

    update() {
        this.player.update(this.cursors);
    }

    addCollisions() {
        this.physics.add.collider(this.player, this.blockedLayer);
        this.physics.add.collider(this.enemies, this.blockedLayer);
        this.physics.add.overlap(this.player, this.enemiesGroup, this.player.enemyCollision.bind(this.player));
        this.physics.add.overlap(this.player, this.portal, this.loadNextLevel.bind(this, false));
        this.physics.add.overlap(this.coinsGroup, this.player, this.coinsGroup.collectCoin.bind(this.coinsGroup));
    }

    resize(width, height) {
        if(width === undefined){
            width = this.sys.game.config.width;
        }
        if(height === undefined){
            height = this.sys.game.config.height;
        }
        
        this.cameras.resize(width, height);
    }

    createMap() {
        // add water background
        this.add.tileSprite(0, 0, 8000, 8000, 'RPGpack_sheet', 31);
        // create the tilemap
        this.map = this.make.tilemap({key: this._LEVELS[this._LEVEL]});
        // add tileset image
        this.tiles = this.map.addTilesetImage('RPGpack_sheet');
        // create our layers
        this.backgroundLayer = this.map.createStaticLayer('Background', this.tiles, 0, 0);
        this.blockedLayer = this.map.createStaticLayer('Blocked', this.tiles, 0, 0);
        this.blockedLayer.setCollisionByExclusion([-1]);
    }

    createPlayer() {
        this.map.findObject('Player', (obj) => {
            if(this._NEWGAME && this._LEVEL == 1){
                if(obj.type == "StartingPosition"){
                    this.player = new Player(this, obj.x, obj.y);
                }
            }else {
                this.player = new Player(this, obj.x, obj.y);
            }
        });
    }

    createPortal() {
        this.map.findObject('Portal', (obj) => {
            if(this._LEVEL === 1){
                this.portal = new Portal(this, obj.x, obj.y-68 );
            }else {
                this.portal = new Portal(this, obj.x, obj.y+68 );
            }
        });
    }

    loadNextLevel (endGame = false) {
        if(!this.loadingLevel){
            this.cameras.main.fade(200, 90, 90, 100);
            this.cameras.main.on('camerafadeoutcomplete', () => {
                if(endGame){
                    this.scene.restart({level: 1, levels: this._LEVELS, newGame: true})
                }else if(this._LEVEL === 1){
                    this.scene.restart({level: 2, levels: this._LEVELS, newGame: false})
                }else if(this._LEVEL === 2){
                    this.scene.restart({level: 1, levels: this._LEVELS, newGame: false})
                }
            });
            this.loadingLevel = true;
        }
    }
}