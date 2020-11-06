const config = {
    rowCount: 10,
    colCount: 10,
    startPositionX: 0,
    startPositionY: 0,
};
const player = {
    x: null,
    y: null,
    step: [],

    init(startX, startY) {
        this.x = startX;
        this.y = startY;
    },
    saveStep(direction) {this.step.push(direction)},
    getStep(stepNumber) {                                                                           //task 2.1
        if(stepNumber <= 0) alert('Нумерация шагов начинается с 1');
        else if(this.step.length === 0) alert('Вы не сделали еще ни одного шага');
        else if(stepNumber > this.step.length - 1) alert(`Вы совершили всего ${this.step.length-1} шагов.`);
        else alert(this.step[stepNumber-1]);
    },
    isTop(y) {return y === 0},                                                                      //task 2.2
    isDown(y) {return y === (config.rowCount - 1)},
    isLeft(x) {return x === 0},
    isRight(x) {return x === (config.colCount - 1)},
    move(direction) {
        switch (direction) {
            case 2:
                if(!this.isDown(this.y)){
                    this.y++;
                    this.saveStep('down');
                }
                break;
            case 4:
                if(!this.isLeft(this.x)){
                    this.x--;
                    this.saveStep('left');
                }
                break;
            case 5:
                const  stepNumber = parseInt(prompt('Введите номер шага:'))
                this.getStep(stepNumber);
                break;
            case 6:
                if(!this.isRight(this.x)){
                    this.x++;
                    this.saveStep('right');
                }
                break;
            case 8:
                if(!this.isTop(this.y)) {
                    this.y--;
                    this.saveStep('up');
                }
                break;
        }
    },
};

const game = {
    settings: config,
    player,

    run() {
        this.player.init(this.settings.startPositionX, this.settings.startPositionY);

        while (true) {
            this.render();

            const direction = this.getDirection();

            if (direction === -1) return alert('До свидания!');

            this.player.move(direction);
        }
    },

    render() {
        let map = '';

        for (let row = 0; row < this.settings.rowCount; row++) {
            for (let col = 0; col < this.settings.colCount; col++) {
                if (this.player.y === row && this.player.x === col) {
                    map += '0 ';
                } else {
                    map += 'X '
                }
            }
            map += '\n';
        }

        console.clear();
        console.log(map);
    },

    getDirection() {
        const availableDirections = [-1, 2, 4, 5, 6, 8];

        while (true) {
            const direction = parseInt(prompt('Введите число куда хотите переместиться, 5 получит ход по номеру, -1 для выхода'));

            if (!availableDirections.includes(direction)) {
                alert(`Для перемещения необходимо ввести одно из чисел: ${availableDirections.join(', ')}.`);
                continue;
            }

            return direction;
        }
    },
}

game.run();
