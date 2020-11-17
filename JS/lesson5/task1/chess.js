'use strict';

function createChessBoard() {
    const table = document.getElementById('chess');
    for(let i = 0; i <=8; i++){
        const tr = document.createElement('tr');
        for(let j = 0; j <=8; j++){
            if(i === 0) {
                const th = document.createElement('th');
                if(j>0){
                    // A = 65
                    let de = new TextDecoder('utf-8');
                    th.innerText = de.decode(new Uint8Array([64 + j])); //преобразуем байт код в строку
                }
                tr.appendChild(th);
            }else{
                if(j===0){
                    const th = document.createElement('th');
                    th.innerText = String(i);
                    tr.appendChild(th);
                } else {
                    const td = document.createElement('td');
                    if((i+j)%2 === 0) td.style.backgroundColor = '#000';
                    else td.style.backgroundColor = '#fff';
                    tr.appendChild(td);
                }
            }
        }
        table.appendChild(tr);
    }
}
function arrangeFigures() {
    const tr = document.getElementById("chess").querySelectorAll('tr');

    for(let i = 0; i <=8; i++){
        for(let j = 0; j <=8; j++){
            if(i === 1 || i == 8) {
                const tdAll = tr[`${i}`].querySelectorAll('td');
                if(j > 0){
                    let color;
                    if(i === 1) color = 'blue';
                    else color = 'cyan';
                    //const td = tdAll[`${j}`];
                    tdAll[`${j-1}`].innerText = 'Т';    tdAll[`${j-1}`].style.color = color; j++;
                    tdAll[`${j-1}`].innerText = 'К';    tdAll[`${j-1}`].style.color = color; j++;
                    tdAll[`${j-1}`].innerText = 'Л';    tdAll[`${j-1}`].style.color = color; j++;
                    tdAll[`${j-1}`].innerText = 'КК';   tdAll[`${j-1}`].style.color = color; j++;
                    tdAll[`${j-1}`].innerText = 'Ф';    tdAll[`${j-1}`].style.color = color; j++;
                    tdAll[`${j-1}`].innerText = 'Л';    tdAll[`${j-1}`].style.color = color; j++;
                    tdAll[`${j-1}`].innerText = 'К';    tdAll[`${j-1}`].style.color = color; j++;
                    tdAll[`${j-1}`].innerText = 'Т';    tdAll[`${j-1}`].style.color = color; j++;
                }
            } else if(i === 2 || i == 7) {
                const tdAll = tr[`${i}`].querySelectorAll('td');
                let color;
                if (j > 0) {
                    if (i === 2) color = 'blue';
                    else color = 'cyan';
                    tdAll[`${j - 1}`].innerText = 'П'; tdAll[`${j-1}`].style.color = color;
                }
            }
        }
    }
}
function chessInit() {
    createChessBoard();
    arrangeFigures();
}

window.onload = chessInit;