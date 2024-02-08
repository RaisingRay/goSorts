package main

Insertion(arr []int){
	flagMove:=false
	for i:=0;i<len(arr);i++{
		if i!=0 && arr[i]>arr[i-1]	{
			for k:=i-1;k>=0;k--{
				if arr[i]< arr[k] {
					flagMove=true;
				} else if flagMove{
					
				} else if !flagMove{
					arr[i],arr[k]=arr[k],arr[i]	
				}
			}
		}
	}
}