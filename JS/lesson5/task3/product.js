'use strict';
const settings = {
    width: '300px',
    backgroundColor: 'peachpuff',
}
const product = {
    id: null,
    name: null,
    description: null,
    price: null,
    imageUrl: 0,
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
const catalogPosition = {
    product: product,
    setProduct(id, name, price, description, imageUrl) {
        this.product = {...product};
        this.product.id = id;
        this.product.name = name;
        this.product.price = price;
        this.product.description = description;
        this.product.imageUrl = imageUrl;
    },
    getHTML() {
        const div = document.createElement('div');  div.style.display = 'inline-block';          div.style.width = '150px'; div.style.margin = '30px';  div.style.textAlign = 'center';
        const image = document.createElement('div'); image.style.height = '150px';    image.style.width = '150px'; image.style.backgroundImage = `url(${this.product.imageUrl})`; image.style.backgroundColor = '#ddd'
        const name = document.createElement('div');
        const description = document.createElement('div');
        const price = document.createElement('div');
        name.innerText = this.product.name;
        description.innerText = this.product.description;
        price.innerText = this.product.price;
        div.appendChild(image);
        div.appendChild(name);
        div.appendChild(description);
        div.appendChild(price);
        return div;
    }
}
const catalog = {
    productList: [],
    image: null,
    addPosition(id, name, price, description, imageUrl) {
        let position = {...catalogPosition};
        position.setProduct(id, name, price, description, imageUrl)
        position.id = id;
        position.name = name;
        position.price = price;
        position.description = description;
        this.productList.push(position);
    },
    setHTML(tagID) {
        const element = document.getElementById(tagID);
        for(let i = 0; i < this.productList.length; i++){
            element.appendChild(this.productList[i].getHTML());
        }
    }
}
const basketT = basket;
basketT.addPosition(1, 'apple', 10.5, 3);
basketT.addPosition(2, 'orange', 11, 5);
basketT.addPosition(3, 'peach', 20.75, 1);
basketT.addPosition(4, 'ananas', 10, 4);

basketT.setHTML('productBasket');

const catalogT = catalog;
catalogT.addPosition(1, 'apple', 10.5, 'какое то описание данного товара', '123.jpg');
catalogT.addPosition(2, 'orange', 11, 'какое то описание данного товара', '123.jpg');
catalogT.addPosition(3, 'peach', 20.75, 'какое то описание данного товара', '123.jpg');
catalogT.addPosition(4, 'ananas', 10, 'какое то описание данного товара', '123.jpg');

catalog.setHTML('productCatalog');

// console.log(basketT);
// console.log(`priceForAll = ${basketT.getPrice()}`);
