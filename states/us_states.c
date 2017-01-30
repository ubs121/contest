#include <stdio.h>
#include <stdlib.h>
#include <string.h>

// дараах бүтэц өгөгдсөн учраас шууд оруулав
typedef struct state
{
    int lengthOfState;
    char* name;
    long int population;
    long int populationOver18;
} State;

// дараах prototype-ууд өгөгдсөн учраас шууд оруулав
void readStates(State**, char*);
void printStateInfo(State*, int);
void findTotalPercentage(State*, float*);
void freeStates(State*);

// line-г задалж state бүтэц руу хувиргах туслах функц
void extractState(char* line, State* state);

/***
*    үндсэн логик
*/
int main(int argc, char *argv[]) {
  // файлын нэр заагаагүй бол програмыг таслах
  if (argc < 2) {
    printf("Input file name is missing ! \n");
    return 1;
  }
  // 51 мужид зориулж санах ой хувиарлах
  State* states;
  states = malloc(sizeof(State) * 51);

  // мужуудыг файлаас унших (argv[1]-д файлын нэр байгаа)
  readStates(&states, argv[1]);

 printf("&states = %p\n", states);

  int choice = 0;
  do {
    printf("Put in a state ID to get information or -1 to quit: ");
    scanf("%d", &choice);

    // quit сонгосон бол гарах
    if (choice == -1) {
      // гарах үед 18-с дээш настны эзлэх хувийг харуулах
      float over18 = 0.0f;
      findTotalPercentage(states, &over18);
      printf("Percentage of US that is 18+: %.2f%%\n", over18);
      break;
    }

    // сонгосон мужийн мэдээллийг хэвлэх
    printStateInfo(states, choice);

  } while (1);

  // санах ойг чөлөөлөх
  freeStates(states);

  // амжилттай дууссан болохыг OS-д мэдэгдэх
  return 0;
}








/***
*
*    Функцүүдийн implementation хэсэг эндээс эхлэнэ
*
*/
void readStates(State** states, char* filename) {
 FILE * fp;
 char * line = NULL;
 size_t len = 0;
 ssize_t read;

 fp = fopen(filename, "r");
 if (fp == NULL) {
     printf("Can't read file: %s\n", filename);
     exit(EXIT_FAILURE);
 }

 State* p = *states;
 while ((read = getline(&line, &len, fp)) != -1) {
     // <Length of the state 1 name>
     p -> lengthOfState = atoi(line);

     // <State 1 name> <State 1 population> <State 1 population over 18>
     if ((read = getline(&line, &len, fp)) != -1) {

       extractState(line, p);
     }

     p++;
 }

 fclose(fp);

 if (line) {
     free(line);
 }
}

void extractState(char* line, State* state) {
  // мужийн нэр
  state -> name = malloc(sizeof(char) * state -> lengthOfState);
  strncpy(state -> name, line, state -> lengthOfState);

  // бусад тоон өгөгдлийн эхэнд заагчийг байрлуулах
  char* p = line;
  p = p +  state -> lengthOfState;

  // population
  while (p != NULL && '0'<= *p && *p <= '9') {
    state -> population = state -> population * 10 + (*p - 48);
    p++;
  }

  p++; // сул зай алгасах

  // populationOver18
  while (p != NULL && '0'<= *p && *p <= '9') {
    state -> populationOver18 = state -> populationOver18 * 10 + (*p - 48);
    p++;
  }
}

void printStateInfo(State* states, int id) {
  printf("------State info------\n");
  printf("State name: %s\n", states[id].name);
  printf("Length of the name: %d\n", states[id].lengthOfState);
  printf("Total population: %ld\n", states[id].population);
  printf("Total population over 18: %ld\n", states[id].populationOver18);
  printf("\n");
}

void findTotalPercentage(State* states, float* f) {
   long int totalPop = 0L;
   long int totalOver18 = 0L;

   for (int i=0; i< 51; i++) {
     totalPop += states[i].population;
     totalOver18 += states[i].populationOver18;
   }

  *f = (double)totalOver18 / (double)totalPop * 100.0;
}

void freeStates(State* states) {
  free(states);
}