'use strict';
const settings = {
    width: '300px',
    backgroundColor: 'peachpuff',
}
const product = {
    id: null,
    name: null,
    description: null,
    price: null
};
const basketPosition = {
    count: null,
    product: product,
    getPriceTotal() {
        return this.count * this.product.price;
    },
    setProduct(id, name, price) {
        this.product = {...product};
        this.product.id = id;
        this.product.name = name;
        this.product.price = price;
    },
    getHTML() {
        const div = document.createElement('div');
        const name = document.createElement('div');         name.style.display = 'inline-block';          name.style.width = '150px';
        const price = document.createElement('div');        price.style.display = 'inline-block';         price.style.width = '50px';         price.style.textAlign = 'center';
        const count = document.createElement('div');        count.style.display = 'inline-block';         count.style.width = '50px';         count.style.textAlign = 'center';
        const totalPrice = document.createElement('div');   totalPrice.style.display = 'inline-block';    totalPrice.style.width = '50px';    totalPrice.style.textAlign = 'center';
        name.innerText = this.product.name;
        price.innerText = this.product.price;
        count.innerText = this.count;
        totalPrice.innerText = this.getPriceTotal();
        div.appendChild(name);
        div.appendChild(price);
        div.appendChild(count);
        div.appendChild(totalPrice);
        return div;
    }
};

const basket = {
    positionList: [],
    getPrice() {
        let priceTotal = 0;
        for(let i = 0; i < this.positionList.length; i++)
            priceTotal += this.positionList[i].getPriceTotal();
        return priceTotal;
    },
    addPosition(id, name, price, count) {
        let position = {...basketPosition};
        position.setProduct(id, name, parseFloat(price));
        position.count = parseInt(count);
        this.positionList.push(position);
    },
    getHeaderHTML() {
        const div = document.createElement('div');
        const name = document.createElement('div');         name.style.display = 'inline-block';          name.style.width = '150px';
        const price = document.createElement('div');        price.style.display = 'inline-block';         price.style.width = '50px';         price.style.textAlign = 'center';
        const count = document.createElement('div');        count.style.display = 'inline-block';         count.style.width = '50px';         count.style.textAlign = 'center';
        const totalPrice = document.createElement('div');   totalPrice.style.display = 'inline-block';    totalPrice.style.width = '50px';    totalPrice.style.textAlign = 'center';
        name.innerText = 'Наименование';
        price.innerText = 'Цена';
        count.innerText = 'Кол.';
        totalPrice.innerText = 'Всего';
        div.appendChild(name);
        div.appendChild(price);
        div.appendChild(count);
        div.appendChild(totalPrice);
        return div;
    },
    getEmptyHTML() {
        const div = document.createElement('div');
        const empty = document.createElement('div');    empty.style.textAlign = 'center';
        empty.innerText = 'Корзина пуста!';
        return div.appendChild(empty);
    },
    getTotalHTML() {
        const div = document.createElement('div');          div.style.backgroundColor = '#ddd';
        const name = document.createElement('div');         name.style.display = 'inline-block';          name.style.width = '250px ';
        const totalPrice = document.createElement('div');   totalPrice.style.display = 'inline-block';    totalPrice.style.width = '50px';    totalPrice.style.textAlign = 'center';
        name.innerText = 'Всего';
        totalPrice.innerText = String(this.getPrice());
        div.appendChild(name);
        div.appendChild(totalPrice);
        return div;
    },
    setHTML(tagID) {
        const element = document.getElementById(tagID);
        element.style.width = settings.width;
        element.style.backgroundColor = settings.backgroundColor;
        if( this.positionList.length === 0 ){
            element.appendChild(this.getEmptyHTML());
        }else {
            element.appendChild(this.getHeaderHTML());
            for(let i = 0; i < this.positionList.length; i++){
                element.appendChild(this.positionList[i].getHTML());
            }
            element.appendChild(this.getTotalHTML());
        }
    }
}

let basketT = basket;
basketT.addPosition(1, 'apple', 10.5, 3);
basketT.addPosition(2, 'orange', 11, 5);
basketT.addPosition(3, 'peach', 20.75, 1);
basketT.addPosition(4, 'ananas', 10, 4);

basketT.setHTML('productBasket');

// console.log(basketT);
// console.log(`priceForAll = ${basketT.getPrice()}`);
