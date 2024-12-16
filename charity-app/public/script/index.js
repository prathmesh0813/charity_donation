const addData = async (event) => {
    event.preventDefault();
    const userName = document.getElementById("userName").value;
    const name = document.getElementById("name").value;
    const address = document.getElementById("address").value;
    const dob = document.getElementById("dob").value;
    const gender = document.getElementById("gender").value;
    const contactNo = document.getElementById("contactNo").value;
    const emailId = document.getElementById("emailId").value;
    const typeofBeans = document.getElementById("typeofBeans").value;
    const color = document.getElementById("color").value;
    const status = document.getElementById("status").value;


    const farmerData = {
        userName: userName,
        name: name,
        address: address,
        dob: dob,
        gender: gender,
        contactNo: contactNo,
        emailId: emailId,
        typeofBeans: typeofBeans,
        color: color,
        status: status,

    }

    if (
        userName.length == 0 ||
        name.length == 0 ||
        address.length == 0 ||
        dob.length == 0 ||
        gender.length == 0 ||
        contactNo.length == 0 ||
        emailId.length == 0 ||
        typeofBeans.length == 0 ||
        color.length == 0 ||
        status.length == 0
    ) {
        alert("Please enter the data properly.");
    } else {
        try {
            const response = await fetch("/api/farmer", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(farmerData)
            })

            return alert("Farmer Registered");
        } catch (err) {
            alert("Error");
            console.log(err);
        }
    }
}

const readData = async (event) => {
    event.preventDefault();
    const userName = document.getElementById("userNameInput").value;

    if (userName.length == 0) {
        alert("Please enter a valid UserName.");
    } else {
        try {
            const response = await fetch(`/api/farmer / ${ userName }`);
            let responseData = await response.json();
            console.log("response data", responseData);
            alert(JSON.stringify(responseData));
        } catch (err) {
            alert("Error");
            console.log(err);
        }
    }
};