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

window.onload = createChessBoard;