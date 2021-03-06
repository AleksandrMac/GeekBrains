"use strict";

const settings = {
    rowsCount: 21,
    colsCount: 21,
    speed: 2,
    winFoodCount: 50,
};

const config = {
    settings,

    init(userSettings = {}) {
        Object.assign(this.settings, userSettings);
    },

    getRowsCount() {
        return this.settings.rowsCount;
    },

    getColsCount() {
        return this.settings.colsCount;
    },

    getSpeed() {
        return this.settings.speed;
    },

    getWinFoodCount() {
        return this.settings.winFoodCount;
    },

    validate() {
        const result = {
            isValid: true,
            errors: [],
        };

        if (this.settings.rowsCount < 10 || this.settings.rowsCount > 30) {
            result.isValid = false;
            result.errors.push('Неверные настройки, значение rowsCount должно быть в диапазоне от 10 до 30 включительно.');
        }

        if (this.settings.colsCount < 10 || this.settings.colsCount > 30) {
            result.isValid = false;
            result.errors.push('Неверные настройки, значение colsCount должно быть в диапазоне от 10 до 30 включительно.');
        }

        if (this.settings.speed < 1 || this.settings.speed > 10) {
            result.isValid = false;
            result.errors.push('Неверные настройки, значение speed должно быть в диапазоне от 1 до 10 включительно.');
        }

        if (this.settings.winFoodCount < 1 || this.settings.winFoodCount > 50) {
            result.isValid = false;
            result.errors.push('Неверные настройки, значение winFoodCount должно быть в диапазоне от 1 до 10 включительно.');
        }

        return result;
    },
};

const map = {
    cells: {},
    header: null,
    usedCells: [],

    init(rowsCount, colsCount) {
        const table = document.getElementById('game');
        table.innerHTML = '';

        this.cells = {} // { x1_y1: td, x1_y2: td, ... , xn_yn}
        this.usedCells = [];

        this.header = document.createElement('h1');
        this.setHeaderText();
        table.insertAdjacentElement("beforebegin", this.header);
        for (let row = 0; row < rowsCount; row++) {
            const tr = document.createElement('tr');
            tr.classList.add('row');
            table.appendChild(tr);

            for (let col = 0; col < colsCount; col++) {
                const td = document.createElement('td');
                td.classList.add('cell');

                this.cells[`x${col}_y${row}`] = td;
                tr.appendChild(td);
            }
        }
    },

    render(snakePointsArray, foodPoint, barrierPoint) {
        for (const cell of this.usedCells) {
            cell.className = 'cell';
        }

        this.setHeaderText(snakePointsArray.length);
        this.usedCells = [];

        snakePointsArray.forEach((point, index) => {
            const snakeCell = this.cells[`x${point.x}_y${point.y}`];
            // console.log(snakeCell);
            snakeCell.classList.add(index === 0 ? 'snakeHead' : 'snakeBody');
            this.usedCells.push(snakeCell);
        });

        const foodCell = this.cells[`x${foodPoint.x}_y${foodPoint.y}`];
        foodCell.classList.add('food');
        this.usedCells.push(foodCell);
        
        if(barrierPoint.x !== null && barrierPoint.y !== null) {
            const barrierCell = this.cells[`x${barrierPoint.x}_y${barrierPoint.y}`];
            barrierCell.classList.add('barrier');
            this.usedCells.push(barrierCell);
        }
    },
    setHeaderText(snakeLength = 1) {
        switch (snakeLength) {
            case 1: 
                this.header.innerHTML = `Для начала игры нажмите кнопку "Старт"`; 
                break;
            default:
                this.header.innerHTML = `Ваш счет: ${snakeLength - 1}`;
        }
    }
};

const snake = {
    body: [],
    direction: null,
    lastStepDirection: null,

    init(startBody, direction) {
        this.body = startBody;
        this.direction = direction;
        this.lastStepDirection = direction;
    },

    getBody() {
        return this.body;
    },

    getLastStepDirection() {
        return this.lastStepDirection;
    },

    isOnPoint(point) {
        return this.body.some(snakePoint => snakePoint.x === point.x && snakePoint.y === point.y);
    },

    makeStep() {
        this.lastStepDirection = this.direction;
        this.body.unshift(this.getNextStepHeadPoint());
        this.body.pop();
    },

    growUp() {
      const lastBodyIndex = this.body.length - 1;
      const lastBodyPoint = this.body[lastBodyIndex];
      const lastBodyPointClone = Object.assign({}, lastBodyPoint);

      this.body.push(lastBodyPointClone);
    },

    getNextStepHeadPoint() {
        let firstPoint = this.body[0];
        /*if(firstPoint.x === 0) firstPoint.x = settings.colsCount;
        if(firstPoint.x === settings.colsCount) firstPoint.x = 0;
        if(firstPoint.y === 0) firstPoint.x = settings.rowsCount;
        if(firstPoint.y === settings.rowsCount) firstPoint.x = 0;*/

        switch (this.direction) {
            case 'up':
                return {x: firstPoint.x, y: (firstPoint.y + settings.rowsCount - 1) % settings.rowsCount};
            case 'right':
                return {x: (firstPoint.x + 1) % settings.colsCount, y: firstPoint.y};
            case 'down':
                return {x: firstPoint.x, y: (firstPoint.y + 1) % settings.rowsCount};
            case 'left':
                return {x: (firstPoint.x + settings.rowsCount - 1) % settings.colsCount, y: firstPoint.y};
        }
    },

    setDirection(direction) {
        this.direction = direction;
    },
};

const food = {
    x: null,
    y: null,

    getCoordinates() {
        return {
            x: this.x,
            y: this.y,
        }
    },

    setCoordinates(point) {
        this.x = point.x;
        this.y = point.y;
    },

    isOnPoint(point) {
        return this.x === point.x && this.y === point.y;
    },
};

const barrier = {
    x: null,
    y: null,

    getCoordinates() {
        return {
            x: this.x,
            y: this.y,
        }
    },

    setCoordinates(point) {
        this.x = point.x;
        this.y = point.y;
    },

    isOnPoint(point) {
        return this.x === point.x && this.y === point.y;
    },
};

const status = {
    condition: null,

    setPlaying() {
        this.condition = 'playing';
    },

    setStopped() {
        this.condition = 'stopped';
    },

    setFinished() {
        this.condition = 'finished';
    },

    isPlaying() {
        return this.condition === 'playing';
    },

    isStopped() {
        return this.condition === 'stopped';
    },
};

const game = {
    config,
    map,
    snake,
    food,
    barrier,
    status,
    tickCount: 0,
    tickInterval: null,

    init(userSettings) {
        this.config.init(userSettings);
        const validation = this.config.validate();

        if (!validation.isValid) {
            for (const err of validation.errors) {
                console.log(err);
            }

            return;
        }

        this.map.init(this.config.getRowsCount(), this.config.getColsCount());

        this.setEventHandlers();
        this.reset();
    },

    reset() {
        this.stop();
        this.snake.init(this.getStartSnakeBody(), 'up');
        this.food.setCoordinates(this.getRandomFreeCoordinates());
        this.render();
    },

    render() {
        this.map.render(this.snake.getBody(), this.food.getCoordinates(), this.barrier.getCoordinates());
    },

    play() {
        this.status.setPlaying();
        this.tickInterval = setInterval(() => this.tickHandler(), 1000 / this.config.getSpeed());
        this.setPlayButtonState('Стоп');
    },

    stop() {
        this.status.setStopped();
        clearInterval(this.tickInterval);
        this.setPlayButtonState('Старт');
    },

    finish() {
        this.status.setFinished();
        clearInterval(this.tickInterval);
        this.setPlayButtonState('Игра окончена', true);
    },

    tickHandler() {
        if (!this.canMakeStep()) {
            /**
             * вот она ошибка, забыли return =(
             */
            return this.finish();
        }

        if (this.food.isOnPoint(this.snake.getNextStepHeadPoint())) {
            this.snake.growUp();
            this.food.setCoordinates(this.getRandomFreeCoordinates());
            

            if (this.isGameWon()) {
                this.finish();
            }
        }this.snake.makeStep();
        this.render();
        this.tickCount++;
        if(this.tickCount % 20 === 0) this.barrier.setCoordinates(this.getRandomFreeCoordinates());
        
    },

    canMakeStep() {
        const nextHeadPoint = this.snake.getNextStepHeadPoint();

        return !this.snake.isOnPoint(nextHeadPoint)
            /*&& nextHeadPoint.x < this.config.getColsCount()
            && nextHeadPoint.y < this.config.getRowsCount()
            && nextHeadPoint.x >= 0
            && nextHeadPoint.y >= 0*/
            && !(nextHeadPoint.x === this.barrier.getCoordinates().x
            && nextHeadPoint.y === this.barrier.getCoordinates().y);
    },

    isGameWon() {
      return this.snake.getBody().length > this.config.getWinFoodCount();
    },

    setPlayButtonState(text, isDisabled = false) {
        const playButton = document.getElementById('playButton');

        playButton.textContent = text;
        isDisabled ? playButton.classList.add('disabled') : playButton.classList.remove('disabled');
    },

    getStartSnakeBody() {
        return [
            {
                x: Math.floor(this.config.getColsCount() / 2),
                y: Math.floor(this.config.getRowsCount() / 2),
            }
        ];
    },

    getRandomFreeCoordinates() {
        const exclude = [this.food.getCoordinates(), this.barrier.getCoordinates(), ...this.snake.getBody()];

        while (true) {
            const rndPoint = {
                x: Math.floor(Math.random() * this.config.getColsCount()),
                y: Math.floor(Math.random() * this.config.getRowsCount()),
            };

            if (!exclude.some(exPoint => rndPoint.x === exPoint.x && rndPoint.y === exPoint.y)) {
                return rndPoint;
            }
        }
    },

    setEventHandlers() {
        document.getElementById('playButton').addEventListener('click', () => {
            this.playClickHandler();
        });
        document.getElementById('newGameButton').addEventListener('click', () => {
            this.newGameClickHandler();
        });
        document.addEventListener('keydown', event => this.keyDownHandler(event));
    },

    playClickHandler() {
        if (this.status.isPlaying()) {
            this.stop();
        } else if (this.status.isStopped()) {
            this.play();
        }
    },

    newGameClickHandler() {
        this.reset();
    },

    keyDownHandler(event) {
        if (!this.status.isPlaying()) return;

        const direction = this.getDirectionByCode(event.code);

        if (this.canSetDirection(direction)) {
            this.snake.setDirection(direction);
        }
    },

    getDirectionByCode(code) {
       switch (code) {
           case 'KeyW':
           case 'ArrowUp':
               return 'up';
           case 'KeyD':
           case 'ArrowRight':
               return 'right';
           case 'KeyS':
           case 'ArrowDown':
               return 'down';
           case 'KeyA':
           case 'ArrowLeft':
               return 'left';
       }
    },

    canSetDirection(direction) {
        const lastStepDirection = this.snake.getLastStepDirection();

        return direction === 'up' && lastStepDirection !== 'down'
            || direction === 'right' && lastStepDirection !== 'left'
            || direction === 'down' && lastStepDirection !== 'up'
            || direction === 'left' && lastStepDirection !== 'right';
    }
};

game.init({speed: 5});
