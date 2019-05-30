import 'phaser';

export default class Player extends Phaser.Physics.Arcade.Sprite {
    constructor(scene, x, y){
        super(scene, x, y, 'characters', 1);
        this.scene = scene;

        // enable physics
        this.scene.physics.world.enable(this);
        // add our player to the scene
        this.scene.add.existing(this);
        // scale our player
        this.setScale(2);
        this.velocity = 150;
    }

    update(cursors) {
        this.setVelocity(0);
        // check if up or down arrow is pressed
        if(cursors.up.isDown){
            this.setVelocityY(-this.velocity);
        }else if (cursors.down.isDown){
            this.setVelocityY(this.velocity);
        }

        // check if left or right arrow is pressed
        if(cursors.left.isDown){
            this.setVelocityX(-this.velocity);
        }else if (cursors.right.isDown){
            this.setVelocityX(this.velocity);
        }
    }
}