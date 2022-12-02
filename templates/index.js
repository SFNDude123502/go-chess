
const moveForm = document.querySelector("#input")

moveFormSubmit = () => {
    const loginData = new FormData(form);
    const st = loginData.get('st');
    const end = loginData.get('end');
    

    const req = new XMLHttpRequest()
    req.open("GET", "/move?st="+st+"&end="+end, false);
    req.send();

    return true;
}

