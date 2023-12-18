const mapSliceFuncM =<I,O>(array:I[],func:(item:I,index:number)=>O[])=>{
    const result:O[] = []
    for(let i=0;i<array.length;i++){
        result.push(...func(array[i],i))
    }
    return result
}

