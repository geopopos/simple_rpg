import 'phaser';

export default class Player extends Phaser.Physics.Arcade.Sprite {
    constructor(scene, x, y){
        super(scene, x, y, 'characters', 1);
        this.scene = scene;
        this.health = 3;
        this.hitDelay = false;

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
        if(cursors.space.isDown){
            this.velocity = 1000;
        }else if(cursors.space.isUp){
            this.velocity = 150;
        }
        // check if up or down arrow is pressed
        if(cursors.up.isDown){
            this.setVelocityY(-this.velocity);
            this.anims.play('up', true);
        }else if (cursors.down.isDown){
            this.setVelocityY(this.velocity);
            this.anims.play('down', true);
        }else if(cursors.left.isDown){
            this.setVelocityX(-this.velocity);
            this.anims.play('left', true);
        }else if (cursors.right.isDown){
            this.setVelocityX(this.velocity);
            this.anims.play('right', true);
        }
    }

    loseHealth() {
        this.health--;
        this.scene.events.emit('loseHealth', this.health);
        if(this.health === 0) {
            this.scene.loadNextLevel(true);
        }
    }

    enemyCollision(player, enemy){
        if(!this.hitDelay){
            this.loseHealth();
            this.hitDelay = true;
            this.tint = 0xff0000;
            this.scene.time.addEvent({
                delay: 1200,
                callback: () => {
                    this.hitDelay = false;
                    this.tint = 0xffffff;
                },
                callbackScope: this
            });
        }
    }
}