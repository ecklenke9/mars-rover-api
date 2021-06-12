<h1>Mars Rover API</h1>
The Mars Rover API returns the last 10 days of Curiosity 
navigation camera images with a limit of 3 images per day.

<h1>Languages and Tools Used</h1>
<p align="left"> <a href="https://golang.org" target="_blank"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/go/go-original.svg" alt="go" width="40" height="40"/> </a> <a href="https://www.mongodb.com/" target="_blank"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/mongodb/mongodb-original-wordmark.svg" alt="mongodb" width="40" height="40"/> </a> </p>

<h1>Installation</h1>
Clone the Mars Rover API repository to your local system:

```sh 
git clone https://github.com/ecklenke9/mars-rover-api.git
```

<h1>Running the Application</h1>
Run the following cmd at the root level of the application to fetch 
Curiosity Rover images:

```sh 
make rover
```

<h1>API Endpoint (Optional)</h1>
This application can also be used as a web api hosted on a server.

Fetch the Curiosity Rover images:

```http request
GET http://localhost:8080/roverImages
```

Output:

```json
{
    "2021-06-02": [
        "https://mars.nasa.gov/msl-raw-images/proj/msl/redops/ods/surface/sol/03136/opgs/edr/fcam/FLB_675895467EDR_F0880804FHAZ00318M_.JPG",
        "https://mars.nasa.gov/msl-raw-images/proj/msl/redops/ods/surface/sol/03136/opgs/edr/fcam/FRB_675895467EDR_F0880804FHAZ00318M_.JPG",
        "https://mars.nasa.gov/msl-raw-images/proj/msl/redops/ods/surface/sol/03136/opgs/edr/fcam/FLB_675895220EDR_D0880798TRAV15030M_.JPG"
    ],
    "2021-06-03": [
        "https://mars.nasa.gov/msl-raw-images/proj/msl/redops/ods/surface/sol/03137/opgs/edr/fcam/FLB_675980024EDR_F0880804FHAZ00341M_.JPG",
        "https://mars.nasa.gov/msl-raw-images/proj/msl/redops/ods/surface/sol/03137/opgs/edr/fcam/FRB_675980024EDR_F0880804FHAZ00341M_.JPG",
        "https://mars.nasa.gov/msl-raw-images/proj/msl/redops/ods/surface/sol/03137/opgs/edr/rcam/RRB_675980058EDR_F0880804RHAZ00341M_.JPG"
    ],
    "2021-06-04": [
        "https://mars.nasa.gov/msl-raw-images/proj/msl/redops/ods/surface/sol/03138/opgs/edr/fcam/FLB_676075087EDR_F0881230FHAZ00302M_.JPG",
        "https://mars.nasa.gov/msl-raw-images/proj/msl/redops/ods/surface/sol/03138/opgs/edr/fcam/FRB_676075087EDR_F0881230FHAZ00302M_.JPG",
        "https://mars.nasa.gov/msl-raw-images/proj/msl/redops/ods/surface/sol/03138/opgs/edr/rcam/RLB_676075208EDR_F0881230RHAZ00311M_.JPG"
    ],
    "2021-06-05": [
        "https://mars.nasa.gov/msl-raw-images/proj/msl/redops/ods/surface/sol/03139/opgs/edr/fcam/FLB_676170830EDR_F0881230FHAZ00337M_.JPG",
        "https://mars.nasa.gov/msl-raw-images/proj/msl/redops/ods/surface/sol/03139/opgs/edr/fcam/FRB_676170830EDR_F0881230FHAZ00337M_.JPG",
        "https://mars.nasa.gov/msl-raw-images/proj/msl/redops/ods/surface/sol/03139/opgs/edr/fcam/FLB_676170699EDR_F0881230FHAZ00206M_.JPG"
    ],
    "2021-06-06": [
        "https://mars.nasa.gov/msl-raw-images/proj/msl/redops/ods/surface/sol/03140/opgs/edr/fcam/FLB_676252514EDR_F0881734FHAZ00302M_.JPG",
        "https://mars.nasa.gov/msl-raw-images/proj/msl/redops/ods/surface/sol/03140/opgs/edr/fcam/FRB_676252514EDR_F0881734FHAZ00302M_.JPG",
        "https://mars.nasa.gov/msl-raw-images/proj/msl/redops/ods/surface/sol/03140/opgs/edr/fcam/FLB_676239866EDR_F0881230FHAZ00236M_.JPG"
    ],
    "2021-06-07": [
        "https://mars.nasa.gov/msl-raw-images/proj/msl/redops/ods/surface/sol/03141/opgs/edr/fcam/FLB_676335807EDR_F0881734FHAZ00341M_.JPG",
        "https://mars.nasa.gov/msl-raw-images/proj/msl/redops/ods/surface/sol/03141/opgs/edr/fcam/FRB_676335807EDR_F0881734FHAZ00341M_.JPG",
        "https://mars.nasa.gov/msl-raw-images/proj/msl/redops/ods/surface/sol/03141/opgs/edr/rcam/RRB_676335841EDR_F0881734RHAZ00341M_.JPG"
    ],
    "2021-06-08": [
        "https://mars.nasa.gov/msl-raw-images/proj/msl/redops/ods/surface/sol/03142/opgs/edr/fcam/FLB_676437875EDR_F0881734FHAZ00337M_.JPG",
        "https://mars.nasa.gov/msl-raw-images/proj/msl/redops/ods/surface/sol/03142/opgs/edr/fcam/FRB_676437875EDR_F0881734FHAZ00337M_.JPG",
        "https://mars.nasa.gov/msl-raw-images/proj/msl/redops/ods/surface/sol/03142/opgs/edr/fcam/FLB_676437744EDR_F0881734FHAZ00206M_.JPG"
    ],
    "2021-06-09": [
        "https://mars.nasa.gov/msl-raw-images/proj/msl/redops/ods/surface/sol/03145/opgs/edr/fcam/FLB_676696218EDR_F0882422FHAZ00302M_.JPG",
        "https://mars.nasa.gov/msl-raw-images/proj/msl/redops/ods/surface/sol/03145/opgs/edr/fcam/FRB_676696218EDR_F0882422FHAZ00302M_.JPG",
        "https://mars.nasa.gov/msl-raw-images/proj/msl/redops/ods/surface/sol/03145/opgs/edr/fcam/FLB_676681070EDR_F0882130FHAZ00236M_.JPG"
    ],
    "2021-06-10": [
        "https://mars.nasa.gov/msl-raw-images/proj/msl/redops/ods/surface/sol/03145/opgs/edr/fcam/FLB_676696218EDR_F0882422FHAZ00302M_.JPG",
        "https://mars.nasa.gov/msl-raw-images/proj/msl/redops/ods/surface/sol/03145/opgs/edr/fcam/FRB_676696218EDR_F0882422FHAZ00302M_.JPG",
        "https://mars.nasa.gov/msl-raw-images/proj/msl/redops/ods/surface/sol/03145/opgs/edr/fcam/FLB_676681070EDR_F0882130FHAZ00236M_.JPG"
    ],
    "2021-06-11": [
        "https://mars.nasa.gov/msl-raw-images/proj/msl/redops/ods/surface/sol/03145/opgs/edr/fcam/FLB_676696218EDR_F0882422FHAZ00302M_.JPG",
        "https://mars.nasa.gov/msl-raw-images/proj/msl/redops/ods/surface/sol/03145/opgs/edr/fcam/FRB_676696218EDR_F0882422FHAZ00302M_.JPG",
        "https://mars.nasa.gov/msl-raw-images/proj/msl/redops/ods/surface/sol/03145/opgs/edr/fcam/FLB_676681070EDR_F0882130FHAZ00236M_.JPG"
    ]
}
```
