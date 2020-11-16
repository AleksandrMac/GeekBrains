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
        return `<div class = "basketItem">
                    <div>${this.product.name}</div>
                    <div>${this.product.price}</div>
                    <div>${this.count}</div>
                    <div>${this.getPriceTotal()}</div>
                </div>`;
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
        return `<div class="basketItem">
                    <div>Наименование</div>
                    <div>Цена</div>
                    <div>Кол.</div>
                    <div>Всего</div>
                </div>`;
    },
    getEmptyHTML() {
        return `<div style="text-align : center">Корзина пуста!</div>`;
    },
    getTotalHTML() {
        return `<div style="background: #ddd">
                    <div style="display : inline-block; width : 250px">Всего</div>
                    <div style="display : inline-block; width : 50px; text-align : center">${String(this.getPrice())}</div>
                </div>`;
    },
    setHTML(tagID) {
        const element = document.getElementById(tagID);
        if( this.positionList.length === 0 ){
            element.insertAdjacentHTML("beforeend", this.getEmptyHTML());
        }else {
            element.insertAdjacentHTML("beforeend", this.getHeaderHTML());
            for(let i = 0; i < this.positionList.length; i++){
                element.insertAdjacentHTML("beforeend", this.positionList[i].getHTML());
            }
            element.insertAdjacentHTML("beforeend", this.getTotalHTML());
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
        return `<div class = "catalogItem">
                    <div style="background-image: url(${this.product.imageUrl})"></div>
                    <div>${this.product.name}</div>
                    <div>${this.product.description}</div>
                    <div>${this.product.price}</div>
                </div>`;
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
            element.insertAdjacentHTML("beforeend", this.productList[i].getHTML());
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