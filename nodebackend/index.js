const express=require('express');
const cors=require('cors');
const morgan=require('morgan');

const app=express();

app.use(cors());
app.use(express.json({limit:'50mb'}));
app.use(express.urlencoded({limit:'50mb',extended:true}));
app.use(morgan('dev'));

app.get('/',(req,res)=>{
    res.status(200).json({message:"GET API is working!"});
})

app.post('/',(req,res)=>{
    res.status(200).json({data:req.body});
})

app.listen(8000,()=>{
    console.log("Server has been started!");
})