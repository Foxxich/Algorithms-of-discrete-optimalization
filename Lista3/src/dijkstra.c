#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <uinterface.h>
#include <heap.h>
#include <list.h>

#define INFINITE  0xFFFFFFFF

void dijkstra(Graph *V, int s, int t, char run);


int main(int argc, char **argv)
{
printf("OK");
    generateResult(argc, argv, &dijkstra);
    return 0;
}

Element createNewElement(Node* nd)
{
    Element result;
    result.e = nd;
    result.value = &nd->value;
    result.index = &nd->position.index;
    return result;
}

void dijkstra(Graph *V, int s, int t, char run)
{
    int i;
    Heap heap;
    initHeap(V->edgesAmount, &heap);
    for(i=0; i<V->nodesAmount; i++)
    {
        Node* actualNode = &V->nodeList[i];
        V->nodeList[i].prev = NULL;
        V->nodeList[i].position.index = -1;
        actualNode->value = INFINITE;
    }
    // printf("%d %d \n", s, t);
    // printf("%d \n",V->nodeList[V->nodesAmount].value);
    V->nodeList[s-1].value = 0;
    push(createNewElement(&V->nodeList[s-1]), &heap);
    Element u;
    ListNode *v;
    while((u = pop(&heap)).e != NULL)
    {
        Node *actualNode = ((Node*)u.e);
        //printf("%d\n", actualNode->value);
        if(t != -1)
            if(actualNode->id == t)
                break;
        for(v = actualNode->neighboursList.listBegin; v; v = v->nextNode)
        {
            Neighbour *actualNeighbour = (Neighbour*)v->value;
            unsigned int neighbourIndex = actualNeighbour->neighbour;
            unsigned int newVal = actualNode->value + actualNeighbour->value;
            //printf("%d\n", actualNeighbour->value);
            if(V->nodeList[neighbourIndex-1].value > newVal)
            { 
                Node *neighbour = &V->nodeList[neighbourIndex-1];
                neighbour->value = newVal;
                if(neighbour->position.index == -1)
                    push(createNewElement(neighbour), &heap);
                if(run == 'p')
                   increaseKey(neighbour->position.index, newVal, &heap);
                //printf("%d", &heap);
                neighbour->prev = u.e;
            }
        }
    }
    releaseHeap(&heap);
}
