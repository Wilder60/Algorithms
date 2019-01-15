using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Sorts
{
    class Program
    {
/*---------------------------------------------------------------------------------------------------------------------*/
        static int[] InsertionSort(int[] Array)
        {
            for (int i = 1; i < Array.Length; i++)
            {
                int key = Array[i];
                int j = i - 1;
                while (j > -1 && Array[j] > key)
                {
                    Array[j + 1] = Array[j];
                    j--;
                }
                Array[j + 1] = key;
            }
            return Array;
        }
/*---------------------------------------------------------------------------------------------------------------------*/
        static int[] BubbleSort(int[] Array)
        {
            for(int i = 0; i < Array.Length -1; i++)
            {
                for(int j = 0; j < Array.Length -i-1; j++)
                {
                    if(Array[j] > Array[j + 1])
                    {
                        swap(ref Array[j], ref Array[j + 1]);
                    }
                }
            }
            return Array;
        }
/*---------------------------------------------------------------------------------------------------------------------*/
        static int[] SelectionSort(int[] Array)
        {
            for(int i = 0; i < Array.Length; i++)
            {
                int lowestIndex = i;
                for(int j = i+1; j < Array.Length; j++)
                {
                    if(Array[lowestIndex] > Array[j])
                    {
                        lowestIndex = j;
                    }
                }
                swap(ref Array[lowestIndex], ref Array[i]);
            }
            return Array;
        }
/*---------------------------------------------------------------------------------------------------------------------*/
        static int[] QuickSort(int[] Array)
        {
            return QuickSort(Array, 0, Array.Length - 1);
        }

        static int[] QuickSort(int[] Array, int low, int high)
        {
            if(low < high)
            {
                int pivot = Parition(Array, low, high);
                QuickSort(Array, low, pivot - 1);
                QuickSort(Array, pivot + 1, high);
            }
            return Array;
        }

        static int Parition( int[] Array, int low, int high)
        {
            int pivot = Array[high];
            int i = low - 1;
            for(int j = low; j <= high -1; j++)
            {
                if (Array[j] <= pivot) {
                    i++;
                    swap(ref Array[i], ref Array[j]);
                }
            }
            swap(ref Array[i + 1], ref Array[high]);
            return i + 1;
        }
/*-------------------------------------------------------------------------------------------------------------------*/
        static void swap(ref int value1, ref int value2)
        {
            int temp = value2;
            value2 = value1;
            value1 = temp;
        }


        static void Main(string[] args)
        {
            int[] Array = { 1, 0, 4, 3, 2, 6, 7, 9, 8, 5 };
            Array = QuickSort(Array);

            foreach(int ele in Array)
            {
                Console.WriteLine(ele);
            }
            Console.ReadKey();
        }
    }
}
