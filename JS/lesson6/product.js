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
    imageUrl: [],
    mainImageID: 0,
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
    setProduct(id, name, price, description, imageUrl, mainImageID = 0) {
        this.product = {...product};
        this.product.id = id;
        this.product.name = name;
        this.product.price = price;
        this.product.description = description;
        this.product.imageUrl = imageUrl;
        this.product.mainImageID = mainImageID;
    },
    getHTML() {
        return `<div class = "catalogItem">
                    <div data-image = "${this.product.id}" style="background-image: url(${String(this.product.imageUrl[0].url).replace('%size%','_min')})"></div>
                    <div>${this.product.name}</div>
                    <div>${this.product.description}</div>
                    <div>Цена: ${this.product.price}</div>
                    <div data-buy = "${this.product.id}">купить</div>
                </div>`;
    }
}
const catalog = {
    catalogElement: document.getElementById('productCatalog'),
    sliderElement: null,
    sliderCurrID: null,
    productList: [],

    addPosition(id, name, price, description, imageUrl) {
        let position = {...catalogPosition};
        position.setProduct(id, name, price, description, imageUrl)
        position.id = id;
        position.name = name;
        position.price = price;
        position.description = description;
        position.imageUrl = imageUrl;
        this.productList.push(position);
    },
    cellClickHandler(event){        
        if (this.isClickByBuyButton(event)){
            basket.basketElement.dispatchEvent(new CustomEvent('buy-click', {detail: 
                {id: event.target.dataset.buy},
                bubbles: true,
            }));
            return;
        }
        if (this.isClickByImage(event)){
            this.openSlider(event.target.dataset.image);
            return;
        }
        if (this.isClickByPreviousImage(event)){
            this.previousSlider();
            return;
        }
        if (this.isClickByNextImage(event)){
            this.nextSlider();
            return;
        }
        if (this.isClickByCloseImage(event)){
            this.closeSlider();
            return;
        }
    },
    createSlider(productID) {
        return `
            <div class="slider">
                <image id="slider">
                <div class="slider_prev"></div>
                <div class="slider_next"></div>
                <div class="close"></div>
            </div>`;
    },
    init() {
        this.setProductList();
        this.render();
        this.initEventHandler();
    },
    initEventHandler() {
        this.catalogElement.addEventListener('click', event => this.cellClickHandler(event));
    },
    isClickByBuyButton(event)  {
        return event.target.dataset.buy !== undefined ;
    },
    isClickByImage(event)  {
        return event.target.dataset.image !== undefined ;
    },
    isClickByCloseImage(event)  {
        if(event.target.classList[0] === 'close')
            this.closeSlider();
    },
    isClickByPreviousImage(event)  {
        if(event.target.classList[0] === 'slider_prev')
            this.setPreviousImageForSlider();
    },
    isClickByNextImage(event)  {
        if(event.target.classList[0] === 'slider_next')
            this.setNextImageForSlider();
    },
    isCorrectClick(event) {
        return this.isClickByCell(event);
    },
    getImageUrlByID(productID,imgID) {
        for (const val of this.productList) 
            if (val.id === +productID) return String(val.imageUrl[imgID].url).replace('%size%','_max');
        return '';
    },
    getProductByID(productID) {
        for (const val of this.productList) 
            if (val.id === +productID) return val;
        return null;
    },
    openSlider(productID) {
        const element = document.getElementsByTagName('html')[0];
        element.insertAdjacentHTML("beforeend", this.createSlider(productID));
        this.sliderElement = document.getElementsByClassName('slider')[0];
        this.sliderElement.addEventListener('click', event => this.cellClickHandler(event));
        this.setImageUrlForSlider(productID, 0);
    },
    closeSlider() {
        const element = document.getElementsByClassName('slider')[0];
        element.remove();
        this.sliderElement = null;
        this.sliderCurrID = null;
    },
    render(tagID) {
        for(let i = 0; i < this.productList.length; i++){
            this.catalogElement.insertAdjacentHTML("beforeend", this.productList[i].getHTML());
        }
    },    
    setPreviousImageForSlider() {
        for(const val of this.productList)
            if(val.productID = this.sliderCurrID.productID){
                const len = val.imageUrl.length;
                if(len === 1) this.setImageUrlForSlider(this.sliderCurrID.productID, this.sliderCurrID.imgID);
                else if(len > 1 && this.sliderCurrID.imgID === 0){
                    this.setImageUrlForSlider(this.sliderCurrID.productID, len-1);
                } else this.setImageUrlForSlider(this.sliderCurrID.productID, this.sliderCurrID.imgID - 1);
                
            }
    },
    setNextImageForSlider() {
        for(const val of this.productList)
            if(val.productID = this.sliderCurrID.productID){
                const len = val.imageUrl.length;
                if(len === 1) this.setImageUrlForSlider(this.sliderCurrID.productID, this.sliderCurrID.imgID);
                else if(len > 1 && this.sliderCurrID.imgID === len - 1){
                    this.setImageUrlForSlider(this.sliderCurrID.productID, 0);
                } else this.setImageUrlForSlider(this.sliderCurrID.productID, this.sliderCurrID.imgID + 1);
                
            }
    },
    setImageUrlForSlider(productID, imgID) {
        const element = document.getElementById("slider");
        element.setAttribute("src", this.getImageUrlByID(productID,imgID));
        this.sliderCurrID = {productID: productID, imgID: imgID};
    },
    setProductList() {
        const test = [
            {
                id: 1, name: 'apple', 
                price: 10.5, 
                description: 'какое то описание данного товара', 
                imageUrl: [
                    {url: 'img/green_apple%size%.jpg'},
                    {url: 'img/apple1%size%.jpg'},
                    {url: 'img/apple2%size%.jpg'}
                ],
                mainImageID: 0
            },{
                id: 2, 
                name: 'orange', 
                price: 11, 
                description: 'какое то описание данного товара', 
                imageUrl: [
                    {url: 'img/green_apple%size%.jpg'},
                    {url: 'img/apple1%size%.jpg'},
                    {url: 'img/apple2%size%.jpg'}
                ], 
                mainImageID: 0
            },{
                id: 3, 
                name: 'peach', 
                price: 20.75, 
                description: 'какое то описание данного товара', 
                imageUrl: [
                    {url: 'img/green_apple%size%.jpg'},
                    {url: 'img/apple1%size%.jpg'},
                    {url: 'img/apple2%size%.jpg'}
                ],
                mainImageID: 0
            },{
                id: 4, name: 'ananas', 
                price: 10, 
                description: 'какое то описание данного товара', 
                imageUrl: [
                    {url: 'img/green_apple%size%.jpg'},
                    {url: 'img/apple1%size%.jpg'},
                    {url: 'img/apple2%size%.jpg'}
                ], 
                mainImageID: 0
            }
        ];
        for (const val of test) this.addPosition(val.id, val.name, val.price, val.description, val.imageUrl, val.mainImageID);
    }
}
basket.init();
catalog.init();