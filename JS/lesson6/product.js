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
    id: null,
    count: 0,
    getPriceTotal() {
        return this.count * catalog.getProductByID(this.id).price;
    },
    getHTML() {
        const product = catalog.getProductByID(this.id);
        return `<div class = "basketItem">
                    <div>${product.name}</div>
                    <div>${product.price}</div>
                    <div>${this.count}</div>
                    <div>${this.getPriceTotal()}</div>
                </div>`;
    }
};

const basket = {
    basketElement: document.getElementById('productBasket'),
    positionList: [],
    getPrice() {
        let priceTotal = 0;
        for(let i = 0; i < this.positionList.length; i++)
            priceTotal += this.positionList[i].getPriceTotal();
        return priceTotal;
    },
    addPosition(productID, count) {

        for (let val of this.positionList){
            if(val.id === productID) {

                val.count += count;
                return;
            }
        }
        let position = {...basketPosition};
        position.id = productID;
        position.count = count;
        this.positionList.push(position);
    },
    contain(productID) {
        for (val of this.positionList) 
            if(val.id === productID) return true;
        return false;
    },
    init() {
        this.render();
        this.initEventHandler();
    },
    initEventHandler(){
        this.basketElement.addEventListener('buy-click', event => { 
            this.addPosition(event.detail.id, 1);
            this.render();
        });
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
    render() {
        //const element = document.getElementById(tagID);
        if( this.positionList.length === 0 ){
            this.basketElement.innerHTML = this.getEmptyHTML();
        }else {
            this.basketElement.innerHTML = this.getHeaderHTML();
            for(let i = 0; i < this.positionList.length; i++){
                this.basketElement.insertAdjacentHTML("beforeend", this.positionList[i].getHTML());
            }
            this.basketElement.insertAdjacentHTML("beforeend", this.getTotalHTML());
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
                    <div>Цена: ${this.product.price}</div>
                    <div data-buy = "${this.product.id}">купить</div>
                </div>`;
    }
}
const catalog = {
    catalogElement: document.getElementById('productCatalog'),
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
    cellClickHandler(event){        
        if (!this.isCorrectClick(event)) return;
        //console.log(event.dataset);
        //basket.addPosition(event.dataset.buy);
        basket.basketElement.dispatchEvent(new CustomEvent('buy-click', {detail: 
            {id: event.target.dataset.buy},
            bubbles: true,
        }));
    },
    init() {
        this.setProductList();
        this.render();
        this.initEventHandler();
    },
    initEventHandler() {
        this.catalogElement.addEventListener('click', event => this.cellClickHandler(event));
    },
    isClickByCell(event)  {
        return event.target.dataset.buy !== undefined ;
    },
    isCorrectClick(event) {
        return this.isClickByCell(event);
    },
    getProductByID(productID) {
        for (const val of this.productList) 
            if (val.id === +productID) return val;
        return null;
    },
    render(tagID) {
        //const element = document.getElementById(tagID);
        for(let i = 0; i < this.productList.length; i++){
            this.catalogElement.insertAdjacentHTML("beforeend", this.productList[i].getHTML());
        }
    },
    setProductList() {
        const test = [
            {id: 1, name: 'apple', price: 10.5, description: 'какое то описание данного товара', imageUrl: '123.jpg'},
            {id: 2, name: 'orange', price: 11, description: 'какое то описание данного товара', imageUrl: '123.jpg'},
            {id: 3, name: 'peach', price: 20.75, description: 'какое то описание данного товара', imageUrl: '123.jpg'},
            {id: 4, name: 'ananas', price: 10, description: 'какое то описание данного товара', imageUrl: '123.jpg'}
        ];
        for (const val of test) this.addPosition(val.id, val.name, val.price, val.description, val.imageUrl);
    }
}
/*const basketT = basket;
basketT.addPosition(1, 'apple', 10.5, 3);
basketT.addPosition(2, 'orange', 11, 5);
basketT.addPosition(3, 'peach', 20.75, 1);
basketT.addPosition(4, 'ananas', 10, 4);*/

basket.init();
catalog.init();

/*const catalogT = catalog;
catalogT.addPosition(1, 'apple', 10.5, 'какое то описание данного товара', '123.jpg');
catalogT.addPosition(2, 'orange', 11, 'какое то описание данного товара', '123.jpg');
catalogT.addPosition(3, 'peach', 20.75, 'какое то описание данного товара', '123.jpg');
catalogT.addPosition(4, 'ananas', 10, 'какое то описание данного товара', '123.jpg');

catalog.init('productCatalog');*/