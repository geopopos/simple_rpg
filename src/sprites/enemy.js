import 'phaser';

export default class Enemy extends Phaser.Physics.Arcade.Sprite {
    constructor(scene, x, y){
        super(scene, x, y, 'characters', 4);
        this.scene = scene;

        // enable physics
        this.scene.physics.world.enable(this);
        // add our enemy to the scene
        this.scene.add.existing(this);
        // scale our enemy
        this.setScale(2);
        

        // move our enemy
        this.timeEvent = this.scene.time.addEvent({
            delay: 3000,
            callback: this.move,
            loop:true,
            callbackScope: this
        });
    }

    move () {
        const randNumber = Math.floor((Math.random() * 4) + 1);
        switch(randNumber){
            case 1:
                this.setVelocityX(100);
                break;
            case 2:
                this.setVelocityX(-100);
                break;
            case 3:
                this.setVelocityY(100);
                break;
            case 4:
                this.setVelocityY(-100);
                break;
            default:
                throw "directional number for enemy movement must be in the range 1-4";
        }
        this.scene.time.addEvent({
            delay: 500,
            callback: () => {
                this.setVelocity(0);
            },
            callbackScope: this
        });
    }
}